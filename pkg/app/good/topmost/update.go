package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if _, err := topmostmwcli.UpdateTopMost(ctx, &topmostmwpb.TopMostReq{
		ID:                     h.ID,
		AppID:                  h.AppID,
		Title:                  h.Title,
		Message:                h.Message,
		Posters:                h.Posters,
		StartAt:                h.StartAt,
		EndAt:                  h.EndAt,
		ThresholdCredits:       h.ThresholdCredits,
		RegisterElapsedSeconds: h.RegisterElapsedSeconds,
		ThresholdPaymentAmount: h.ThresholdPaymentAmount,
		KycMust:                h.KycMust,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
