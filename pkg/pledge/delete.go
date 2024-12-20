package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
)

func (h *Handler) DeletePledge(ctx context.Context) (*npool.Pledge, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkPledge(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetPledge(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := pledgemwcli.DeletePledge(ctx, h.ID, h.EntID, h.GoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
