package pledge

import (
	"context"

	apppledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
	apppledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

// TODO: check start mode with power rental start mode

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdatePledge(ctx context.Context) (*npool.AppPledge, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPledge(ctx); err != nil {
		return nil, err
	}
	if err := apppledgemwcli.UpdatePledge(ctx, &apppledgemwpb.PledgeReq{
		ID:                  h.ID,
		EntID:               h.EntID,
		AppGoodID:           h.AppGoodID,
		Purchasable:         h.Purchasable,
		EnableProductPage:   h.EnableProductPage,
		ProductPage:         h.ProductPage,
		Online:              h.Online,
		Visible:             h.Visible,
		Name:                h.Name,
		DisplayIndex:        h.DisplayIndex,
		Banner:              h.Banner,
		ServiceStartAt:      h.ServiceStartAt,
		EnableSetCommission: h.EnableSetCommission,
		StartMode:           h.StartMode,
	}); err != nil {
		return nil, err
	}
	return h.GetPledge(ctx)
}
