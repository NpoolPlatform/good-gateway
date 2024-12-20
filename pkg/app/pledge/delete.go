package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	apppledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePledge(ctx context.Context) (*npool.AppPledge, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPledge(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetPledge(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid pledge")
	}
	if err := apppledgemwcli.DeletePledge(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
