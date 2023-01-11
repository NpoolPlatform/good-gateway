//nolint:nolintlint
package migrator

import (
	"context"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
)

func Migrate(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
