package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

	"github.com/google/uuid"
)

func (h *Handler) CreateTopMost(ctx context.Context) (*npool.TopMost, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := topmostmwcli.CreateTopMost(ctx, &topmostmwpb.TopMostReq{
		EntID:                  h.EntID,
		AppID:                  h.AppID,
		TopMostType:            h.TopMostType,
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
