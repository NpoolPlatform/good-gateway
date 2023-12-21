package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/google/uuid"
)

func (h *Handler) CreateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	// TODO: check exist of topmost and appgood

	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := topmostgoodmwcli.CreateTopMostGood(ctx, &topmostgoodmwpb.TopMostGoodReq{
		EntID:     h.EntID,
		AppGoodID: h.AppGoodID,
		TopMostID: h.TopMostID,
		Posters:   h.Posters,
		Price:     h.Price,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
