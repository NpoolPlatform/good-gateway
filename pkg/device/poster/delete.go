package poster

import (
	"context"

	postermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/poster"
	postermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePoster(ctx context.Context) (*postermwpb.Poster, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, err
	}

	if err := postermwcli.DeletePoster(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
