package description

import (
	"context"

	appgooddescriptionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/description"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDescription(ctx context.Context) (*npool.Description, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDescription(ctx); err != nil {
		return nil, err
	}

	if err := appgooddescriptionmwcli.DeleteDescription(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetDescription(ctx)
}
