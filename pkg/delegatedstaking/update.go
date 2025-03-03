package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/delegatedstaking"
	delegatedstakingmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
)

func (h *Handler) UpdateDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := delegatedstakingmwcli.UpdateDelegatedStaking(ctx, &delegatedstakingmwpb.DelegatedStakingReq{
		ID:                   h.ID,
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		ContractCodeURL:      h.ContractCodeURL,
		ContractCodeBranch:   h.ContractCodeBranch,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetDelegatedStaking(ctx)
}
