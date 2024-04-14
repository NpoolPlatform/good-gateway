package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDefault(ctx context.Context) (*npool.Default, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDefault(ctx); err != nil {
		return nil, err
	}

	if err := defaultmwcli.DeleteDefault(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetDefault(ctx)
}
