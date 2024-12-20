package pledge

import (
	"context"
	"fmt"

	apppledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/pledge"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	apppledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPledge(ctx context.Context) error {
	exist, err := apppledgemwcli.ExistPledgeConds(ctx, &apppledgemwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid apppledge")
	}
	return nil
}
