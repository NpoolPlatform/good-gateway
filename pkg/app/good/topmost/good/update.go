package topmostgood

import (
	"context"
	"fmt"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (h *Handler) UpdateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	info, err := topmostgoodmwcli.GetTopMostGoodOnly(ctx, &topmostgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid topmostgood")
	}

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
