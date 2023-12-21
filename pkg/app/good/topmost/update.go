package topmost

import (
	"context"
	"fmt"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	info, err := topmostmwcli.GetTopMostOnly(ctx, &topmostmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid topmost")
	}

	if _, err := topmostmwcli.UpdateTopMost(ctx, &topmostmwpb.TopMostReq{
		ID:                     h.ID,
		AppID:                  h.AppID,
		Title:                  h.Title,
		Message:                h.Message,
		Posters:                h.Posters,
		StartAt:                h.StartAt,
		EndAt:                  h.EndAt,
		ThresholdCredits:       h.ThresholdCredits,
		ThresholdPurchases:     h.ThresholdPurchases,
		RegisterElapsedSeconds: h.RegisterElapsedSeconds,
		ThresholdPaymentAmount: h.ThresholdPaymentAmount,
		KycMust:                h.KycMust,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
