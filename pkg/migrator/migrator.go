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
	"github.com/shopspring/decimal"

	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgooddescription "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	entappgooddisplaycolor "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	entappgooddisplayname "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplayname"
	entappgoodlabel "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodlabel"
	entappgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"
	entapplegacypowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/applegacypowerrental"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entdevicemanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	entgoodcoinreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoinreward"
	// entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	// entgoodrewardhistory "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodrewardhistory"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"

	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
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

	rows, err := tx.QueryContext(ctx, "select ent_id,descriptions,created_at,updated_at from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type Description struct {
		AppGoodID    uuid.UUID `json:"ent_id"`
		Descriptions string
		CreatedAt    uint32
		UpdatedAt    uint32
	}
	for rows.Next() {
		des := &Description{}
		if err := rows.Scan(&des.AppGoodID, &des.Descriptions, &des.CreatedAt, &des.UpdatedAt); err != nil {
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
				SetAppGoodID(des.AppGoodID).
				SetDescription(description).
				SetCreatedAt(des.CreatedAt).
				SetUpdatedAt(des.UpdatedAt).
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

	rows, err := tx.QueryContext(ctx, "select ent_id,display_colors,created_at,updated_at from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type DisplayColor struct {
		AppGoodID     uuid.UUID `json:"ent_id"`
		DisplayColors string
		CreatedAt     uint32
		UpdatedAt     uint32
	}
	for rows.Next() {
		displayColor := &DisplayColor{}
		if err := rows.Scan(&displayColor.AppGoodID, &displayColor.DisplayColors, &displayColor.CreatedAt, &displayColor.UpdatedAt); err != nil {
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
				SetAppGoodID(displayColor.AppGoodID).
				SetColor(color).
				SetIndex(uint8(idx)).
				SetCreatedAt(displayColor.CreatedAt).
				SetUpdatedAt(displayColor.UpdatedAt).
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

	rows, err := tx.QueryContext(ctx, "select ent_id,display_names,created_at,updated_at from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type DisplayName struct {
		AppGoodID    uuid.UUID `json:"ent_id"`
		DisplayNames string
		CreatedAt    uint32
		UpdatedAt    uint32
	}
	for rows.Next() {
		displayName := &DisplayName{}
		if err := rows.Scan(&displayName.AppGoodID, &displayName.DisplayNames, &displayName.CreatedAt, &displayName.UpdatedAt); err != nil {
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
				SetAppGoodID(displayName.AppGoodID).
				SetName(name).
				SetIndex(uint8(idx)).
				SetCreatedAt(displayName.CreatedAt).
				SetUpdatedAt(displayName.UpdatedAt).
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

	rows, err := tx.QueryContext(ctx, "select app_good_id,posters,created_at,updated_at from extra_infos where deleted_at = 0")
	if err != nil {
		return err
	}

	type Poster struct {
		AppGoodID uuid.UUID
		Posters   string
		CreatedAt uint32
		UpdatedAt uint32
	}
	for rows.Next() {
		poster := &Poster{}
		if err := rows.Scan(&poster.AppGoodID, &poster.Posters, &poster.CreatedAt, &poster.UpdatedAt); err != nil {
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
				SetCreatedAt(poster.CreatedAt).
				SetUpdatedAt(poster.UpdatedAt).
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

	rows, err := tx.QueryContext(ctx, "select app_good_id,labels,created_at,updated_at from extra_infos where deleted_at = 0")
	if err != nil {
		return err
	}

	type Label struct {
		AppGoodID uuid.UUID
		Labels    string
		CreatedAt uint32
		UpdatedAt uint32
	}
	for rows.Next() {
		label := &Label{}
		if err := rows.Scan(&label.AppGoodID, &label.Labels, &label.CreatedAt, &label.UpdatedAt); err != nil {
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
				SetCreatedAt(label.CreatedAt).
				SetUpdatedAt(label.UpdatedAt).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateDeviceInfo(ctx context.Context, tx *ent.Tx) error {
	facturers, err := tx.DeviceManufacturer.Query().Where(entdevicemanufacturer.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(facturers) > 0 {
		logger.Sugar().Warnw("device manufacturers is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select id,manufacturer,created_at,updated_at from device_infos where deleted_at = 0")
	if err != nil {
		return err
	}

	type Manufacturer struct {
		ID           uint32
		Manufacturer string
		CreatedAt    uint32
		UpdatedAt    uint32
	}

	manufacturerMap := map[string]uuid.UUID{}
	for rows.Next() {
		manufacturer := &Manufacturer{}
		if err := rows.Scan(&manufacturer.ID, &manufacturer.Manufacturer, &manufacturer.CreatedAt, &manufacturer.UpdatedAt); err != nil {
			return err
		}
		manufacturerID, ok := manufacturerMap[manufacturer.Manufacturer]
		if !ok {
			manufacturerID = uuid.New()
			if _, err := tx.
				DeviceManufacturer.
				Create().
				SetName(manufacturer.Manufacturer).
				SetLogo("").
				SetCreatedAt(manufacturer.CreatedAt).
				SetUpdatedAt(manufacturer.UpdatedAt).
				Save(ctx); err != nil {
				return err
			}
			manufacturerMap[manufacturer.Manufacturer] = manufacturerID
		}
		if _, err := tx.
			DeviceInfo.
			UpdateOneID(manufacturer.ID).
			SetManufacturerID(manufacturerID).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func migrateTechnicalFeeRatio(ctx context.Context, tx *ent.Tx) error {
	names, err := tx.AppLegacyPowerRental.Query().Where(entapplegacypowerrental.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	if len(names) > 0 {
		logger.Sugar().Warnw("applegacypowerrentals is not empty")
		return nil
	}

	rows, err := tx.QueryContext(ctx, "select ent_id,technical_fee_ratio,created_at,updated_at from app_goods where deleted_at = 0")
	if err != nil {
		return err
	}

	type TechniqueFee struct {
		AppGoodID         uuid.UUID `json:"ent_id"`
		TechniqueFeeRatio decimal.Decimal
		CreatedAt         uint32
		UpdatedAt         uint32
	}
	for rows.Next() {
		technique := &TechniqueFee{}
		if err := rows.Scan(&technique.AppGoodID, &technique.TechniqueFeeRatio, &technique.CreatedAt, &technique.UpdatedAt); err != nil {
			return err
		}
		if _, err := tx.
			AppLegacyPowerRental.
			Create().
			SetAppGoodID(technique.AppGoodID).
			SetTechniqueFeeRatio(technique.TechniqueFeeRatio).
			SetCreatedAt(technique.CreatedAt).
			SetUpdatedAt(technique.UpdatedAt).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

// after migrate posters & labels
func createExtraInfoForAppGoods(ctx context.Context, tx *ent.Tx) error {
	appgoods, err := tx.AppGood.Query().Where(entappgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	for _, appgood := range appgoods {
		if _, err := tx.
			ExtraInfo.
			Create().
			SetAppGoodID(appgood.EntID).
			SetLikes(0).
			SetDislikes(0).
			SetScore(decimal.NewFromInt(0)).
			SetScoreCount(0).
			SetCommentCount(0).
			SetRecommendCount(0).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func fillAppGoodIDForAppStockLocks(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.AppStock.Query().Where(entappstock.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	appstocks := map[string]uuid.UUID{}
	for _, info := range infos {
		appstocks[info.EntID.String()] = info.AppGoodID
	}

	rows, err := tx.QueryContext(ctx, "select id,app_stock_id from app_stock_locks where app_good_id is null and deleted_at = 0")
	if err != nil {
		return err
	}
	type AppStockLock struct {
		ID         uint32
		AppStockID uuid.UUID
	}
	for rows.Next() {
		lock := &AppStockLock{}
		if err := rows.Scan(&lock.ID, &lock.AppStockID); err != nil {
			return err
		}
		appGoodID, ok := appstocks[lock.AppStockID.String()]
		if !ok {
			appGoodID = uuid.Nil
		}
		if _, err := tx.
			AppStockLock.
			UpdateOneID(lock.ID).
			SetAppGoodID(appGoodID).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func migrateGoods(ctx context.Context, tx *ent.Tx) error {
	goods, err := tx.Good.Query().Where(entgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}

	for _, good := range goods {
		exist, err := tx.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(good.EntID),
				entgoodbase.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := tx.
				GoodBase.
				Create().
				SetEntID(good.EntID).
				SetName(good.Title).
				SetGoodType(good.GoodType).
				SetBenefitType(good.BenefitType).
				SetServiceStartAt(good.StartAt).
				SetStartMode(good.StartMode).
				SetTestOnly(good.TestOnly).
				SetBenefitIntervalHours(good.BenefitIntervalHours).
				SetPurchasable(true).
				SetOnline(true).
				Save(ctx); err != nil {
				return err
			}
		}
		exist, err = tx.
			PowerRental.
			Query().
			Where(
				entpowerrental.GoodID(good.EntID),
				entpowerrental.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := tx.
				PowerRental.
				Create().
				// SetEntID(good.EntID).
				SetGoodID(good.EntID).
				SetDeviceTypeID(good.DeviceInfoID).
				SetVendorLocationID(good.VendorLocationID).
				SetUnitPrice(good.UnitPrice).
				SetQuantityUnit(good.QuantityUnit).
				SetQuantityUnitAmount(good.QuantityUnitAmount).
				SetDeliveryAt(good.DeliveryAt).
				SetUnitLockDeposit(good.UnitLockDeposit).
				SetDurationDisplayType(goodtypes.GoodDurationType_GoodDurationByDay.String()).
				SetStockMode(goodtypes.GoodStockMode_GoodStockByUnique.String()).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateAppGoods(ctx context.Context, tx *ent.Tx) error {
	appgoods, err := tx.AppGood.Query().Where(entappgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}

	for _, appgood := range appgoods {
		exist, err := tx.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(appgood.EntID),
				entappgoodbase.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := tx.
				AppGoodBase.
				Create().
				SetEntID(appgood.EntID).
				SetAppID(appgood.AppID).
				SetGoodID(appgood.GoodID).
				SetPurchasable(appgood.EnablePurchase).
				SetEnableProductPage(appgood.EnableProductPage).
				SetProductPage(appgood.ProductPage).
				SetName(appgood.GoodName).
				SetDisplayIndex(appgood.DisplayIndex).
				SetBanner(appgood.GoodBanner).
				SetOnline(appgood.Online).
				SetVisible(appgood.Visible).
				Save(ctx); err != nil {
				return err
			}
		}
		exist, err = tx.
			AppPowerRental.
			Query().
			Where(
				entapppowerrental.AppGoodID(appgood.EntID),
				entapppowerrental.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := tx.
				AppPowerRental.
				Create().
				// SetEntID(good.EntID).
				SetAppGoodID(appgood.EntID).
				SetServiceStartAt(appgood.ServiceStartAt).
				SetCancelMode(appgood.CancelMode).
				SetCancelableBeforeStartSeconds(appgood.CancellableBeforeStart).
				SetEnableSetCommission(appgood.EnableSetCommission).
				SetMinOrderAmount(appgood.MinOrderAmount).
				SetMaxOrderAmount(appgood.MaxOrderAmount).
				SetMaxUserAmount(appgood.MaxUserAmount).
				SetUnitPrice(appgood.UnitPrice).
				SetSaleStartAt(appgood.SaleStartAt).
				SetSaleEndAt(appgood.SaleEndAt).
				SetSaleMode(goodtypes.GoodSaleMode_GoodSaleModeMainnetSpot.String()).
				SetFixedDuration(true).
				SetPackageWithRequireds(appgood.PackageWithRequireds).
				SetMinOrderDurationSeconds(appgood.MinOrderDuration * 24 * 60 * 60).
				SetMaxOrderDurationSeconds(appgood.MaxOrderDuration * 24 * 60 * 60).
				SetStartMode(goodtypes.GoodStartMode_GoodStartModeNextDay.String()).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateGoodCoins(ctx context.Context, tx *ent.Tx) error {
	goods, err := tx.Good.Query().Where(entgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}

	for _, good := range goods {
		exist, err := tx.
			GoodCoin.
			Query().
			Where(
				entgoodcoin.GoodID(good.EntID),
				entgoodcoin.CoinTypeID(good.CoinTypeID),
				entgoodcoin.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := tx.
				GoodCoin.
				Create().
				SetGoodID(good.EntID).
				SetCoinTypeID(good.CoinTypeID).
				SetMain(true).
				SetIndex(0).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateGoodRewards(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.Good.Query().Where(entgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	goods := map[string]uuid.UUID{}
	for _, info := range infos {
		goods[info.EntID.String()] = info.CoinTypeID
	}
	rows, err := tx.QueryContext(ctx, "select id,ent_id,good_id,reward_state,last_reward_at,reward_tid,next_reward_start_amount,last_reward_amount,last_unit_reward_amount,total_reward_amount,created_at,updated_at from good_rewards where deleted_at = 0")
	if err != nil {
		return err
	}

	type Reward struct {
		ID                    uuid.UUID
		EntID                 uuid.UUID
		GoodID                uuid.UUID
		RewardState           string
		LastRewardAt          uint32
		RewardTid             uuid.UUID
		NextRewardStartAmount decimal.Decimal
		LastRewardAmount      decimal.Decimal
		LastUnitRewardAmount  decimal.Decimal
		TotalRewardAmount     decimal.Decimal
		CreatedAt             uint32
		UpdatedAt             uint32
	}
	for rows.Next() {
		reward := &Reward{}
		if err := rows.Scan(&reward.ID, &reward.EntID, &reward.GoodID, &reward.RewardState, &reward.LastRewardAt, &reward.RewardTid, &reward.LastRewardAmount, &reward.LastUnitRewardAmount, &reward.TotalRewardAmount, &reward.CreatedAt, &reward.UpdatedAt); err != nil {
			return err
		}
		coinTypeID, ok := goods[reward.GoodID.String()]
		if !ok {
			coinTypeID = uuid.Nil
		}

		exist, err := tx.
			GoodCoinReward.
			Query().
			Where(
				entgoodcoinreward.GoodID(reward.GoodID),
				entgoodcoinreward.CoinTypeID(coinTypeID),
				entgoodcoinreward.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}

		if !exist {
			if _, err := tx.
				GoodCoinReward.
				Create().
				SetEntID(reward.EntID).
				SetGoodID(reward.GoodID).
				SetCoinTypeID(coinTypeID).
				SetRewardTid(reward.RewardTid).
				SetNextRewardStartAmount(reward.NextRewardStartAmount).
				SetLastRewardAmount(reward.LastRewardAmount).
				SetLastUnitRewardAmount(reward.LastUnitRewardAmount).
				SetTotalRewardAmount(reward.TotalRewardAmount).
				SetCreatedAt(reward.CreatedAt).
				SetUpdatedAt(reward.UpdatedAt).
				Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func fillCoinTypeIDInGoodRewardHistories(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.Good.Query().Where(entgood.DeletedAt(0)).All(ctx)
	if err != nil {
		return err
	}
	goods := map[string]uuid.UUID{}
	for _, info := range infos {
		goods[info.EntID.String()] = info.CoinTypeID
	}

	rows, err := tx.QueryContext(ctx, "select id,good_id from good_reward_histories where coin_type_id is null and deleted_at = 0")
	if err != nil {
		return err
	}

	type RewardHistory struct {
		ID     uint32
		GoodID uuid.UUID
	}
	for rows.Next() {
		reward := &RewardHistory{}
		if err := rows.Scan(&reward.ID, &reward.GoodID); err != nil {
			return err
		}
		coinTypeID, ok := goods[reward.GoodID.String()]
		if !ok {
			coinTypeID = uuid.Nil
		}
		if _, err := tx.
			GoodRewardHistory.
			UpdateOneID(reward.ID).
			SetCoinTypeID(coinTypeID).
			Save(ctx); err != nil {
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

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		// if err := migrateDescriptions(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateDisplayColors(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateDisplayNames(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migratePosters(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateLabels(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateDeviceInfo(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateTechnicalFeeRatio(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := createExtraInfoForAppGoods(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := fillAppGoodIDForAppStockLocks(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateGoods(ctx, tx); err != nil {
		// 	return err
		// }
		// if err := migrateAppGoods(ctx, tx); err != nil {
		// 	return err
		// }
		return nil
	})
}
