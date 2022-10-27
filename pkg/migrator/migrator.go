//nolint:nolintlint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	"github.com/NpoolPlatform/good-manager/pkg/db"

	cgoodent "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	cconst "github.com/NpoolPlatform/cloud-hashing-goods/pkg/message/const"

	goodpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	sconst "github.com/NpoolPlatform/stock-manager/pkg/message/const"

	stockent "github.com/NpoolPlatform/stock-manager/pkg/db/ent"

	_ "github.com/NpoolPlatform/stock-manager/pkg/db/ent/runtime"
)

func Migrate(ctx context.Context) error {
	err := migrationCloudGoods(ctx)
	if err != nil {
		return err
	}
	return migrationStock(ctx)
}

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 10
	maxIdle      = 10
	MaxLife      = 3
	priceScale12 = 1000000000000
)

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsb", "error", err)
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

//nolint
func migrationCloudGoods(ctx context.Context) (err error) {
	cli, err := db.Client()
	if err != nil {
		return err
	}

	appGood, err := cli.AppGood.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}

	if len(appGood) != 0 {
		return nil
	}

	cloudGood, err := open(cconst.ServiceName)
	if err != nil {
		return err
	}

	defer cloudGood.Close()

	cloudGoodCli := cgoodent.NewClient(cgoodent.Driver(entsql.OpenDB(dialect.MySQL, cloudGood)))

	logger.Sugar().Infow("Migrate goods", "Start", "...")

	defer func() {
		logger.Sugar().Infow("Migrate goods", "Done", "...", "error", err)
	}()

	// AppGood
	appGoodInfos, err := cloudGoodCli.
		AppGood.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	goodInfos, err := cloudGoodCli.
		GoodInfo.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	goodMap := map[uuid.UUID]*cgoodent.GoodInfo{}
	for _, good := range goodInfos {
		goodMap[good.ID] = good
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppGoodCreate, len(appGoodInfos))
		for i, info := range appGoodInfos {
			good, ok := goodMap[info.GoodID]
			goodName := ""
			if ok {
				goodName = good.Title
			}

			bulk[i] = tx.AppGood.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreateAt).
				SetUpdatedAt(info.UpdateAt).
				SetDeletedAt(info.DeleteAt).
				SetAppID(info.AppID).
				SetGoodID(info.GoodID).
				SetOnline(info.Online).
				SetVisible(info.Visible).
				SetGoodName(goodName).
				SetPrice(decimal.NewFromInt(int64(info.Price)).Div(decimal.NewFromInt(priceScale12))).
				SetDisplayIndex(int32(info.DisplayIndex)).
				SetPurchaseLimit(info.PurchaseLimit).
				SetCommissionPercent(int32(info.CommissionPercent))
		}
		_, err = tx.AppGood.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// Comment
	goodCommentInfos, err := cloudGoodCli.
		GoodComment.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.CommentCreate, len(goodCommentInfos))
		for i, info := range goodCommentInfos {
			bulk[i] = tx.Comment.
				Create().
				SetID(info.ID).
				SetCreatedAt(uint32(info.CreateAt / 1e9)).
				SetUpdatedAt(uint32(info.UpdateAt / 1e9)).
				SetDeletedAt(uint32(info.DeleteAt / 1e9)).
				SetAppID(info.AppID).
				SetUserID(info.UserID).
				SetGoodID(info.GoodID).
				SetOrderID(info.OrderID).
				SetContent(info.Content).
				SetReplyToID(info.ReplyToID)
		}
		_, err = tx.Comment.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// DeviceInfo
	deviceInfos, err := cloudGoodCli.
		DeviceInfo.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.DeviceInfoCreate, len(deviceInfos))
		for i, info := range deviceInfos {
			bulk[i] = tx.DeviceInfo.
				Create().
				SetID(info.ID).
				SetCreatedAt(uint32(info.CreateAt / 1e9)).
				SetUpdatedAt(uint32(info.UpdateAt / 1e9)).
				SetDeletedAt(uint32(info.DeleteAt / 1e9)).
				SetType(info.Type).
				SetManufacturer(info.Manufacturer).
				SetPowerComsuption(uint32(info.PowerComsuption)).
				SetShipmentAt(uint32(info.ShipmentAt))
		}
		_, err = tx.DeviceInfo.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// ExtraInfo
	extraInfos, err := cloudGoodCli.
		GoodExtraInfo.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.ExtraInfoCreate, len(extraInfos))
		for i, info := range extraInfos {
			bulk[i] = tx.ExtraInfo.
				Create().
				SetID(info.ID).
				SetCreatedAt(uint32(info.CreateAt / 1e9)).
				SetUpdatedAt(uint32(info.UpdateAt / 1e9)).
				SetDeletedAt(uint32(info.DeleteAt / 1e9)).
				SetGoodID(info.GoodID).
				SetPosters(info.Posters).
				SetLabels(info.Labels).
				SetVoteCount(info.VoteCount).
				SetRating(info.Rating)
		}
		_, err = tx.ExtraInfo.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// Good
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.GoodCreate, len(goodInfos))
		for i, info := range goodInfos {
			goodType := goodpb.GoodType_GoodTypeClassicMining.String()
			if !info.Classic {
				goodType = goodpb.GoodType_GoodTypeUnionMining.String()
			}
			bulk[i] = tx.Good.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreateAt).
				SetUpdatedAt(info.UpdateAt).
				SetDeletedAt(info.DeleteAt).
				SetDeviceInfoID(info.DeviceInfoID).
				SetDurationDays(info.DurationDays).
				SetCoinTypeID(info.CoinInfoID).
				SetInheritFromGoodID(info.InheritFromGoodID).
				SetVendorLocationID(info.VendorLocationID).
				SetPrice(decimal.NewFromInt(int64(info.Price)).Div(decimal.NewFromInt(priceScale12))).
				SetBenefitType(getBenefitType(info.BenefitType.String())).
				SetGoodType(goodType).
				SetTitle(info.Title).
				SetUnit(info.Unit).
				SetUnitAmount(info.UnitPower).
				SetSupportCoinTypeIds(info.SupportCoinTypeIds).
				SetDeliveryAt(info.DeliveryAt).
				SetStartAt(info.StartAt).
				SetTestOnly(false)
		}
		_, err = tx.Good.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// Promotion
	promotionInfos, err := cloudGoodCli.
		AppGoodPromotion.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.PromotionCreate, len(promotionInfos))
		for i, info := range promotionInfos {
			bulk[i] = tx.Promotion.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreateAt).
				SetUpdatedAt(info.UpdateAt).
				SetDeletedAt(info.DeleteAt).
				SetAppID(info.AppID).
				SetGoodID(info.GoodID).
				SetMessage(info.Message).
				SetStartAt(info.Start).
				SetEndAt(info.End).
				SetPrice(decimal.NewFromInt(int64(info.Price)).Div(decimal.NewFromInt(priceScale12)))
		}
		_, err = tx.Promotion.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// Promotion
	recommendInfos, err := cloudGoodCli.
		Recommend.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.RecommendCreate, len(recommendInfos))
		for i, info := range recommendInfos {
			bulk[i] = tx.Recommend.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreateAt).
				SetUpdatedAt(info.UpdateAt).
				SetDeletedAt(info.DeleteAt).
				SetAppID(info.AppID).
				SetGoodID(info.GoodID).
				SetRecommenderID(info.RecommenderID).
				SetMessage(info.Message)
		}
		_, err = tx.Recommend.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	// Promotion
	vendorLocationInfos, err := cloudGoodCli.
		VendorLocation.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.VendorLocationCreate, len(vendorLocationInfos))
		for i, info := range vendorLocationInfos {
			bulk[i] = tx.VendorLocation.
				Create().
				SetID(info.ID).
				SetCreatedAt(uint32(info.CreateAt / 1e9)).
				SetUpdatedAt(uint32(info.UpdateAt / 1e9)).
				SetDeletedAt(uint32(info.DeleteAt / 1e9)).
				SetCountry(info.Country).
				SetProvince(info.Province).
				SetCity(info.City).
				SetAddress(info.Address)
		}
		_, err = tx.VendorLocation.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

//nolint
func migrationStock(ctx context.Context) (err error) {
	cli, err := db.Client()
	if err != nil {
		return err
	}

	stock, err := cli.Stock.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}

	if len(stock) != 0 {
		return nil
	}

	stockManager, err := open(sconst.ServiceName)
	if err != nil {
		return err
	}

	defer stockManager.Close()

	stockManagerCli := stockent.NewClient(stockent.Driver(entsql.OpenDB(dialect.MySQL, stockManager)))

	logger.Sugar().Infow("Migrate stock", "Start", "...")
	defer func() {
		logger.Sugar().Infow("Migrate stock", "Done", "...", "error", err)
	}()

	// AppGood
	stockInfos, err := stockManagerCli.
		Stock.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.StockCreate, len(stockInfos))
		for i, info := range stockInfos {
			bulk[i] = tx.Stock.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreatedAt).
				SetUpdatedAt(info.UpdatedAt).
				SetDeletedAt(info.DeletedAt).
				SetGoodID(info.GoodID).
				SetTotal(info.Total).
				SetLocked(info.Locked).
				SetInService(info.InService).
				SetSold(info.Sold)
		}
		_, err = tx.Stock.CreateBulk(bulk...).Save(_ctx)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func getBenefitType(bType string) string {
	switch bType {
	case "pool":
		return goodpb.BenefitType_BenefitTypePool.String()
	case "platform":
		return goodpb.BenefitType_BenefitTypePlatform.String()
	default:
		return ""
	}
}
