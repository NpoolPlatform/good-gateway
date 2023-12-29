package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	servicename "github.com/NpoolPlatform/good-gateway/pkg/servicename"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 5
	maxIdle      = 2
	MaxLife      = 0
	keyServiceID = "serviceid"
)

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Warnw("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

func migrateGoodOrder(ctx context.Context, conn *sql.DB) error {
	// Get goods
	type good struct {
		EntID        uuid.UUID
		DurationDays uint32
	}
	rows, err := conn.QueryContext(
		ctx,
		"select ent_id,duration_days from good_manager.goods",
	)
	if err != nil {
		return err
	}
	goods := map[uuid.UUID]*good{}
	for rows.Next() {
		g := &good{}
		if err := rows.Scan(&g.EntID, &g.DurationDays); err != nil {
			return err
		}
		goods[g.EntID] = g
	}
	_, err = conn.ExecContext(
		ctx,
		"update good_manager.app_goods set user_purchase_limit='0' where user_purchase_limit is NULL",
	)
	if err != nil {
		return err
	}
	// Get app goods
	type appGood struct {
		EntID             uuid.UUID
		PurchaseLimit     int32
		UserPurchaseLimit decimal.Decimal
		GoodID            uuid.UUID
		DurationDays      uint32
	}
	rows, err = conn.QueryContext(
		ctx,
		"select ent_id,purchase_limit,user_purchase_limit,good_id from good_manager.app_goods",
	)
	if err != nil {
		return err
	}
	appGoods := map[uuid.UUID]*appGood{}
	for rows.Next() {
		ag := &appGood{
			DurationDays: 365,
		}
		if err := rows.Scan(&ag.EntID, &ag.PurchaseLimit, &ag.UserPurchaseLimit, &ag.GoodID); err != nil {
			return err
		}
		if g, ok := goods[ag.GoodID]; ok {
			ag.DurationDays = g.DurationDays
		}
		appGoods[ag.EntID] = ag
	}
	// Update orders
	for appGoodID, ag := range appGoods {
		result, err := conn.ExecContext(
			ctx,
			fmt.Sprintf(
				"update order_manager.orders set duration=%v where app_good_id='%v' and duration=0",
				ag.DurationDays,
				appGoodID,
			),
		)
		if err != nil {
			return err
		}
		_, err = result.RowsAffected()
		if err != nil {
			return err
		}
		result, err = conn.ExecContext(
			ctx,
			fmt.Sprintf(
				"update good_manager.app_goods set min_order_amount='0.1',max_order_amount='%v',max_user_amount='%v',min_order_duration='%v',max_order_duration='%v' where ent_id='%v'",
				ag.PurchaseLimit,
				ag.UserPurchaseLimit,
				ag.DurationDays,
				ag.DurationDays,
				appGoodID,
			),
		)
		if err != nil {
			return err
		}
		_, err = result.RowsAffected()
		if err != nil {
			return err
		}
	}
	return nil
}

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func Migrate(ctx context.Context) error {
	var err error

	logger.Sugar().Infow("Migrate good", "Start", "...")
	defer func() {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Infow("Migrate good", "Done", "...", "error", err)
	}()

	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}

	conn, err := open(servicename.ServiceDomain)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Sugar().Errorw("Close", "Error", err)
		}
	}()

	return migrateGoodOrder(ctx, conn)
}
