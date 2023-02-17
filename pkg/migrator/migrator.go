//nolint:nolintlint
package migrator

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

const LockKey = "stock_migration_lock"

//nolint:funlen
func Migrate(ctx context.Context) error {
	if err := db.Init(); err != nil {
		return err
	}

	cli, err := db.Client()
	if err != nil {
		return err
	}
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

	type stockStruct struct {
		ID        uuid.UUID
		CreatedAt uint32
		UpdatedAt uint32
		DeletedAt uint32
		GoodID    uuid.UUID
		Total     uint32
		Locked    uint32
		InService uint32
		WaitStart uint32
		Sold      uint32
	}

	stocks := []stockStruct{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stockRow, err := tx.QueryContext(ctx, "select * from stocks")
		if err != nil {
			return err
		}

		for stockRow.Next() {
			stock := stockStruct{}
			err = stockRow.Scan(
				&stock.ID,
				&stock.CreatedAt,
				&stock.UpdatedAt,
				&stock.DeletedAt,
				&stock.GoodID,
				&stock.Total,
				&stock.Locked,
				&stock.InService,
				&stock.WaitStart,
				&stock.Sold,
			)
			if err != nil {
				return err
			}
			stocks = append(stocks, stock)
		}

		bulk := make([]*ent.StockV1Create, len(stocks))
		for i, info := range stocks {
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
