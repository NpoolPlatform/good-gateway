package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetPoster(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid poster")
	}

	if err := appgoodpostermwcli.DeletePoster(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
