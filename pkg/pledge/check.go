package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	pledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPledge(ctx context.Context) error {
	exist, err := pledgemwcli.ExistPledgeConds(ctx, &pledgemwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid pledge")
	}
	return nil
}
