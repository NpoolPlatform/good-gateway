package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (h *Handler) UpdateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	// TODO: check exist of topmost and appgood
	if _, err := topmostgoodmwcli.UpdateTopMostGood(ctx, &topmostgoodmwpb.TopMostGoodReq{
		ID:      h.ID,
		Posters: h.Posters,
		Price:   h.Price,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
