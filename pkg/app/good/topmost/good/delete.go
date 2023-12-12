package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (h *Handler) DeleteTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	info, err := topmostgoodmwcli.GetTopMostGoodOnly(ctx, &topmostgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	// TODO: check exist of topmost and appgood
	if _, err := topmostgoodmwcli.DeleteTopMostGood(ctx, *h.ID); err != nil {
		return nil, err
	}

	return h.GetTopMostGoodExt(ctx, info)
}
