package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
)

func (h *Handler) DeleteTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	info, err := h.GetTopMostGood(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: check exist of topmost and appgood
	if _, err := topmostgoodmwcli.DeleteTopMostGood(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
