package displayname

import (
	"context"

	appgooddisplaynamemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/name"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/name"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayName(ctx); err != nil {
		return nil, err
	}

	if err := appgooddisplaynamemwcli.DeleteDisplayName(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return h.GetDisplayName(ctx)
}
