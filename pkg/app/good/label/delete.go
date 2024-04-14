package label

import (
	"context"

	appgoodlabelmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/label"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/label"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteLabel(ctx context.Context) (*npool.Label, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkLabel(ctx); err != nil {
		return nil, err
	}

	if err := appgoodlabelmwcli.DeleteLabel(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetLabel(ctx)
}
