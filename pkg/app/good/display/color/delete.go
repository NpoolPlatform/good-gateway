package displaycolor

import (
	"context"

	appgooddisplaycolormwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/color"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/color"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayColor(ctx); err != nil {
		return nil, err
	}

	if err := appgooddisplaycolormwcli.DeleteDisplayColor(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetDisplayColor(ctx)
}
