package appfee

import (
	"context"

	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

func (h *Handler) UpdateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if err := appfeemwcli.UpdateFee(ctx, &appfeemwpb.FeeReq{
		ID:               h.ID,
		EntID:            h.EntID,
		AppGoodID:        h.AppGoodID,
		ProductPage:      h.ProductPage,
		Name:             h.Name,
		Banner:           h.Banner,
		UnitValue:        h.UnitValue,
		MinOrderDuration: h.MinOrderDuration,
	}); err != nil {
		return nil, err
	}
	return h.GetAppFee(ctx)
}
