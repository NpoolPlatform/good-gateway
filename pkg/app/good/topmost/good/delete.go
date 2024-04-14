package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMostGood(ctx); err != nil {
		return nil, err
	}
	if err := topmostgoodmwcli.DeleteTopMostGood(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetTopMostGood(ctx)
}
