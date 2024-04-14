package description

import (
	"context"

	appgooddescriptionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/description"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
	appgooddescriptionmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDescription(ctx context.Context) (*npool.Description, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDescription(ctx); err != nil {
		return nil, err
	}

	if err := appgooddescriptionmwcli.UpdateDescription(ctx, &appgooddescriptionmwpb.DescriptionReq{
		ID:          h.ID,
		EntID:       h.EntID,
		AppGoodID:   h.AppGoodID,
		Description: h.Description,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDescription(ctx)
}
