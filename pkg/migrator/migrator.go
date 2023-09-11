//nolint:nolintlint
package migrator

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	servicename "github.com/NpoolPlatform/good-gateway/pkg/servicename"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const keyServiceID = "serviceid"

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceName, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func migrateGoodReward(ctx context.Context, tx *ent.Tx) error {
	r, err := tx.QueryContext(ctx, "select id,deleted_at,benefit_state,last_benefit_at,next_benefit_start_amount,last_benefit_amount from goods")
	if err != nil {
		return err
	}
	type g struct {
		ID                     uuid.UUID
		DeletedAt              uint32
		BenefitState           string
		LastBenefitAt          uint32
		NextBenefitStartAmount decimal.Decimal
		LastBenefitAmount      decimal.Decimal
		TotalRewardAmount      decimal.Decimal
	}
	goods := []*g{}
	for r.Next() {
		good := &g{}
		if err := r.Scan(&good.ID, &good.DeletedAt, &good.BenefitState, &good.LastBenefitAt, &good.NextBenefitStartAmount, &good.LastBenefitAmount); err != nil {
			return err
		}
		goods = append(goods, good)
	}

	r, err = tx.QueryContext(ctx, "select good_id,amount from ledger_manager.mining_generals")
	if err != nil {
		return err
	}
	type m struct {
		GoodID uuid.UUID
		Amount decimal.Decimal
	}
	for r.Next() {
		_m := &m{}
		if err := r.Scan(&_m.GoodID, &_m.Amount); err != nil {
			return err
		}
		for _, g := range goods {
			if g.ID == _m.GoodID && g.TotalRewardAmount.Cmp(decimal.NewFromInt(0)) == 0 {
				g.TotalRewardAmount = _m.Amount
			}
		}
	}

	for _, good := range goods {
		reward, err := tx.
			GoodReward.
			Query().
			Where(
				entgoodreward.GoodID(good.ID),
				entgoodreward.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if reward != nil {
			continue
		}
		if _, err := tx.
			GoodReward.
			Create().
			SetGoodID(good.ID).
			SetRewardState(good.BenefitState).
			SetLastRewardAt(good.LastBenefitAt).
			SetNextRewardStartAmount(good.NextBenefitStartAmount).
			SetLastRewardAmount(good.LastBenefitAmount).
			SetTotalRewardAmount(good.TotalRewardAmount).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func migrateAppGoodStock(ctx context.Context, tx *ent.Tx) error {
	stocks, err := tx.
		Stock.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	stockMap := map[uuid.UUID]*ent.Stock{}
	for _, stock := range stocks {
		stockMap[stock.GoodID] = stock
	}
	for _, stock := range stocks {
		if stock.SpotQuantity.Cmp(decimal.NewFromInt(0)) > 0 {
			continue
		}
		if _, err := tx.
			Stock.
			UpdateOneID(stock.ID).
			SetSpotQuantity(stock.Total.Sub(stock.InService).Sub(stock.WaitStart).Sub(stock.Locked)).
			Save(ctx); err != nil {
			return err
		}
	}

	appGoods, err := tx.
		AppGood.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	for _, appGood := range appGoods {
		exist, err := tx.
			AppStock.
			Query().
			Where(
				entappstock.GoodID(appGood.GoodID),
				entappstock.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		stock, ok := stockMap[appGood.GoodID]
		if !ok {
			continue
		}
		if _, err := tx.
			AppStock.
			Create().
			SetAppID(appGood.AppID).
			SetGoodID(appGood.GoodID).
			SetAppGoodID(appGood.ID).
			SetReserved(decimal.NewFromInt(0)).
			SetSpotQuantity(decimal.NewFromInt(0)).
			SetLocked(stock.Locked).
			SetInService(stock.InService).
			SetWaitStart(stock.WaitStart).
			SetSold(stock.Sold).
			Save(ctx); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//nolint
func Migrate(ctx context.Context) error {
	var err error

	if err := db.Init(); err != nil {
		return err
	}
	logger.Sugar().Infow("Migrate order", "Start", "...")
	defer func() {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Infow("Migrate order", "Done", "...", "error", err)
	}()

	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		_, err := tx.
			ExecContext(
				ctx,
				"update extra_infos set score='0' where score is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update good_rewards set total_reward_amount='0' where total_reward_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update good_rewards set last_reward_amount='0' where last_reward_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update good_rewards set last_unit_reward_amount='0' where last_unit_reward_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update good_rewards set next_reward_start_amount='0' where next_reward_start_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update goods set last_benefit_amount='0' where last_benefit_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update goods set next_benefit_start_amount='0' where next_benefit_start_amount is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update goods set unit_lock_deposit='0' where unit_lock_deposit is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update goods set good_type='PowerRenting' where good_type not in ('PowerRenting','MachineRenting','MachineHosting','TechniqueServiceFee','ElectricityFee')",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update app_goods set user_purchase_limit='100000' where user_purchase_limit is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update app_goods set display_colors='[]' where display_colors is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"alter table app_goods modify column technical_fee_ratio decimal(37,18)",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"alter table app_goods modify column electricity_fee_ratio decimal(37,18)",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update stocks_v1 set app_reserved='0' where app_reserved is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"update stocks_v1 set spot_quantity='0' where spot_quantity is NULL",
			)
		if err != nil {
			return err
		}
		_, err = tx.
			ExecContext(
				ctx,
				"alter table recommends modify column recommend_index decimal(37,18)",
			)
		if err != nil {
			return err
		}
		if err := migrateGoodReward(_ctx, tx); err != nil {
			return err
		}
		if err := migrateAppGoodStock(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
