//nolint:nolintlint
package migrator

import (
	"context"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	"github.com/shopspring/decimal"
)

func Migrate(ctx context.Context) error {
	if err := db.Init(); err != nil {
		return err
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.
			AppGood.
			Update().
			SetCommissionSettleType(commmgrpb.SettleType_GoodOrderPercent.String()).
			SetTechnicalFeeRatio(20).
			Save(_ctx)
		if err != nil {
			return err
		}

		_, err = cli.
			Good.
			Update().
			SetNextBenefitStartAmount(decimal.NewFromInt(0)).
			SetLastBenefitAmount(decimal.NewFromInt(0)).
			Save(_ctx)
		if err != nil {
			return err
		}

		infos, err := cli.
			Stock.
			Query().
			All(_ctx)
		if err != nil {
			return err
		}

		for _, info := range infos {
			if info.WaitStart == 0 {
				continue
			}

			_, err := cli.
				Stock.
				UpdateOneID(info.ID).
				SetInService(0).
				SetWaitStart(info.InService).
				Save(_ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
