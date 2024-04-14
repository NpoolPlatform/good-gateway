package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteTopMost(ctx context.Context) (*npool.TopMost, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMost(ctx); err != nil {
		return nil, err
	}

	if err := topmostmwcli.DeleteTopMost(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
