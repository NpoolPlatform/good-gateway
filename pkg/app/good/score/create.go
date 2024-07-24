package score

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/score"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
}

func (h *Handler) CreateScore(ctx context.Context) (*npool.Score, error) {
	handler := &createHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}

	if err := h.CheckUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if h.Score != nil {
		if err := handler.validateScore(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	if err := scoremwcli.CreateScore(ctx, &scoremwpb.ScoreReq{
		EntID:     h.EntID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Score:     h.Score,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetScore(ctx)
}
