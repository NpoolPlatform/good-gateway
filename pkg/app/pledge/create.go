package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	apppledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/pledge"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
	apppledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
	pledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"

	"github.com/google/uuid"
)

type CreateHander struct {
	*Handler
	pledge *pledgemwpb.Pledge
}

func (h *CreateHander) getPledge(ctx context.Context) (err error) {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	h.pledge, err = pledgemwcli.GetPledge(ctx, *h.GoodID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

// TODO: check start mode with power rental start mode
func (h *Handler) CreatePledge(ctx context.Context) (*npool.AppPledge, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	createH := &CreateHander{Handler: h}

	if err := createH.getPledge(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := apppledgemwcli.CreatePledge(ctx, &apppledgemwpb.PledgeReq{
		EntID:               h.EntID,
		AppID:               h.AppID,
		GoodID:              h.GoodID,
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
