package appfee

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteAppFee(ctx context.Context) (*npool.AppFee, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAppFee(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetAppFee(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fee")
	}
	if err := appfeemwcli.DeleteFee(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
