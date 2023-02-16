//nolint:nolintlint
package migrator

import (
	"context"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/shopspring/decimal"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
)

const LockKey = "stock_migration_lock"

func Migrate(ctx context.Context) error {
	if err := db.Init(); err != nil {
		return err
	}

	cli, err := db.Client()
	row, err := cli.StockV1.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}
	if len(row) > 0 {
		return nil
	}

	logger.Sugar().Infow("Migrate order", "Start", "...")
	defer func() {
		_ = redis2.Unlock(LockKey)
		logger.Sugar().Infow("Migrate order", "Done", "...", "error", err)
	}()

	err = redis2.TryLock(LockKey, 0)
	if err != nil {
		return err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		infos, err := tx.
			Stock.
			Query().
			All(_ctx)
		if err != nil {
			return err
		}

		bulk := make([]*ent.StockV1Create, len(infos))
		for i, info := range infos {
			total := decimal.NewFromInt32(int32(info.Total))
			locked := decimal.NewFromInt32(int32(info.Locked))
			inService := decimal.NewFromInt32(int32(info.InService))
			waitStart := decimal.NewFromInt32(int32(info.WaitStart))
			sold := decimal.NewFromInt32(int32(info.Sold))
			bulk[i] = tx.
				StockV1.
				Create().
				SetID(info.ID).
				SetCreatedAt(info.CreatedAt).
				SetUpdatedAt(info.UpdatedAt).
				SetDeletedAt(info.DeletedAt).
				SetGoodID(info.GoodID).
				SetTotal(total).
				SetLocked(locked).
				SetInService(inService).
				SetWaitStart(waitStart).
				SetSold(sold)
			if err != nil {
				return err
			}
		}
		_, err = tx.StockV1.CreateBulk(bulk...).Save(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		_ = redis2.Unlock(LockKey)
		return err
	}

	return nil
}
