package topmost

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetTopMost(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid topmost")
	}

	if err := topmostmwcli.DeleteTopMost(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
