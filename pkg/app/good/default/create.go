package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
	defaultmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *Handler) CreateDefault(ctx context.Context) (*npool.Default, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := defaultmwcli.CreateDefault(ctx, &defaultmwpb.DefaultReq{
		EntID:     h.EntID,
		AppGoodID: h.AppGoodID,
	}); err != nil {
		return nil, err
	}
	return h.GetDefault(ctx)
}
