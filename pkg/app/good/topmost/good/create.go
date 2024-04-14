package topmostgood

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost"
	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/google/uuid"
)

func (h *Handler) CreateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	if err := topmost1.CheckTopMost(ctx, *h.AppID, *h.TopMostID); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := topmostgoodmwcli.CreateTopMostGood(ctx, &topmostgoodmwpb.TopMostGoodReq{
		EntID:        h.EntID,
		AppGoodID:    h.AppGoodID,
		TopMostID:    h.TopMostID,
		UnitPrice:    h.UnitPrice,
		DisplayIndex: h.DisplayIndex,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
