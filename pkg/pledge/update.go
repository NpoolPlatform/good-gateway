package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
	pledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
)

func (h *Handler) UpdatePledge(ctx context.Context) (*npool.Pledge, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkPledge(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := pledgemwcli.UpdatePledge(ctx, &pledgemwpb.PledgeReq{
		ID:                   h.ID,
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		GoodType:             h.GoodType,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		ContractCodeURL:      h.ContractCodeURL,
		ContractCodeBranch:   h.ContractCodeBranch,
		ContractState:        h.ContractState,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetPledge(ctx)
}
