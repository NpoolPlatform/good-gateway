package poster

import (
	"context"

	appgoodpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/poster"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePoster(ctx context.Context) (*npool.Poster, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, err
	}

	if err := appgoodpostermwcli.DeletePoster(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
