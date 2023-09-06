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
	}
	goods := []*g{}
	for r.Next() {
		good := &g{}
		if err := r.Scan(&good.ID, &good.DeletedAt, &good.BenefitState, &good.LastBenefitAt, &good.NextBenefitStartAmount, &good.LastBenefitAmount); err != nil {
			return err
		}
		goods = append(goods, good)
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
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

//nolint:funlen
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
		return nil
	})
}
