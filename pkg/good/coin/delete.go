package goodcoin

import (
	"context"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteGoodCoin(ctx context.Context) (*goodcoinmwpb.GoodCoin, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkGoodCoin(ctx); err != nil {
		return nil, err
	}
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return nil, err
	}
	if err := goodcoinmwcli.DeleteGoodCoin(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, err
}
