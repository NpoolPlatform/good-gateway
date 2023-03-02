//nolint:nolintlint
package migrator

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	constant "github.com/NpoolPlatform/good-gateway/pkg/message/const"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

const keyServiceID = "serviceid"

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(constant.ServiceName, keyServiceID)
	return fmt.Sprintf("migrator:%v", serviceID)
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
		return nil
	})
}
