package appfee

import (
	"context"

	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

func (h *Handler) CreateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if err := appfeemwcli.CreateFee(ctx, &appfeemwpb.FeeReq{
		AppID:                   h.AppID,
		GoodID:                  h.GoodID,
		ProductPage:             h.ProductPage,
		Name:                    h.Name,
		Banner:                  h.Banner,
		UnitValue:               h.UnitValue,
		MinOrderDurationSeconds: h.MinOrderDurationSeconds,
	}); err != nil {
		return nil, err
	}
	return h.GetAppFee(ctx)
}
