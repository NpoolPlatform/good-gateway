package score

import (
	"context"
	"fmt"

	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/score"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateScore(ctx context.Context) (*npool.Score, error) {
	if err := h.CheckUser(ctx); err != nil {
		return nil, err
	}
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if h.Score != nil {
		maxScore := decimal.RequireFromString("10.0")
		score, err := decimal.NewFromString(*h.Score)
		if err != nil {
			return nil, err
		}
		if score.Cmp(maxScore) > 0 {
			return nil, fmt.Errorf("invalid score")
		}
	}

	if err := scoremwcli.CreateScore(ctx, &scoremwpb.ScoreReq{
		EntID:     h.EntID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Score:     h.Score,
	}); err != nil {
		return nil, err
	}

	return h.GetScore(ctx)
}
