//nolint
package migrator

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	servicename "github.com/NpoolPlatform/good-gateway/pkg/servicename"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	entappgooddescription "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	entappgooddisplaycolor "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	entappgooddisplayname "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplayname"
	entappgoodlabel "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodlabel"
	entappgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
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

func migrateDescriptions(ctx context.Context, tx *ent.Tx) error {
	descriptions, err := tx.AppGoodDescription.Query().Where(entappgooddescription.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(descriptions) > 0 {
		logger.Sugar().Warnw("appgooddescriptions is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select ent_id,descriptions from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type Description struct {
		EntID        uuid.UUID
		Descriptions string
	}
	for rows.Next() {
		des := &Description{}
		if err := rows.Scan(&des.EntID, &des.Descriptions); err != nil {
			return err
		}
		var descriptions []string
		if err := json.Unmarshal([]byte(des.Descriptions), &descriptions); err != nil {
			return err
		}
		for idx, description := range descriptions {
			if _, err := tx.
				AppGoodDescription.
				Create().
				SetAppGoodID(des.EntID).
				SetDescription(description).
				SetIndex(uint8(idx)).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateDisplayColors(ctx context.Context, tx *ent.Tx) error {
	colors, err := tx.AppGoodDisplayColor.Query().Where(entappgooddisplaycolor.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(colors) > 0 {
		logger.Sugar().Warnw("appgooddisplaycolors is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select ent_id,display_colors from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type DisplayColor struct {
		EntID         uuid.UUID
		DisplayColors string
	}
	for rows.Next() {
		displayColor := &DisplayColor{}
		if err := rows.Scan(&displayColor.EntID, &displayColor.DisplayColors); err != nil {
			return err
		}
		var colors []string
		if err := json.Unmarshal([]byte(displayColor.DisplayColors), &colors); err != nil {
			return err
		}
		for idx, color := range colors {
			if _, err := tx.
				AppGoodDisplayColor.
				Create().
				SetAppGoodID(displayColor.EntID).
				SetColor(color).
				SetIndex(uint8(idx)).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateDisplayNames(ctx context.Context, tx *ent.Tx) error {
	names, err := tx.AppGoodDisplayName.Query().Where(entappgooddisplayname.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(names) > 0 {
		logger.Sugar().Warnw("appgooddisplaynames is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select ent_id,display_names from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type DisplayName struct {
		EntID        uuid.UUID
		DisplayNames string
	}
	for rows.Next() {
		displayName := &DisplayName{}
		if err := rows.Scan(&displayName.EntID, &displayName.DisplayNames); err != nil {
			return err
		}
		var names []string
		if err := json.Unmarshal([]byte(displayName.DisplayNames), &names); err != nil {
			return err
		}
		for idx, name := range names {
			if _, err := tx.
				AppGoodDisplayName.
				Create().
				SetAppGoodID(displayName.EntID).
				SetName(name).
				SetIndex(uint8(idx)).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migratePosters(ctx context.Context, tx *ent.Tx) error {
	posters, err := tx.AppGoodPoster.Query().Where(entappgoodposter.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(posters) > 0 {
		logger.Sugar().Warnw("appgoodposters is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select app_good_id,posters from extra_infos where deleted_at = 0")
	if err != nil {
		return err
	}

	type Poster struct {
		AppGoodID uuid.UUID
		Posters   string
	}
	for rows.Next() {
		poster := &Poster{}
		if err := rows.Scan(&poster.AppGoodID, &poster.Posters); err != nil {
			return err
		}
		var posters []string
		if err := json.Unmarshal([]byte(poster.Posters), &posters); err != nil {
			return err
		}
		for idx, pos := range posters {
			if _, err := tx.
				AppGoodPoster.
				Create().
				SetAppGoodID(poster.AppGoodID).
				SetPoster(pos).
				SetIndex(uint8(idx)).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateLabels(ctx context.Context, tx *ent.Tx) error {
	labels, err := tx.AppGoodLabel.Query().Where(entappgoodlabel.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(labels) > 0 {
		logger.Sugar().Warnw("appgoodlabels is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select app_good_id,labels from extra_infos where deleted_at = 0")
	if err != nil {
		return err
	}

	type Poster struct {
		AppGoodID uuid.UUID
		Labels    string
	}
	for rows.Next() {
		label := &Poster{}
		if err := rows.Scan(&label.AppGoodID, &label.Labels); err != nil {
			return err
		}
		var labels []string
		if err := json.Unmarshal([]byte(label.Labels), &labels); err != nil {
			return err
		}
		for idx, _label := range labels {
			if _, err := tx.
				AppGoodLabel.
				Create().
				SetAppGoodID(label.AppGoodID).
				SetLabel(_label).
				SetIndex(uint8(idx)).
				Save(ctx); err != nil {
				return err
			}
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

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := migrateDescriptions(ctx, tx); err != nil {
			return err
		}
		if err := migrateDisplayColors(ctx, tx); err != nil {
			return err
		}
		if err := migrateDisplayNames(ctx, tx); err != nil {
			return err
		}
		if err := migratePosters(ctx, tx); err != nil {
			return err
		}
		if err := migrateLabels(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
