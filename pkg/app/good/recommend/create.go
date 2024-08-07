package recommend

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateRecommend(ctx context.Context) (*npool.Recommend, error) {
	if err := h.CheckUserWithUserID(ctx, *h.RecommenderID); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if h.RecommendIndex != nil {
		maxRecommendIndex := decimal.RequireFromString("10.0")
		score, err := decimal.NewFromString(*h.RecommendIndex)
		if err != nil {
			return nil, err
		}
		if score.GreaterThan(maxRecommendIndex) || score.LessThan(decimal.NewFromInt(0)) {
			return nil, wlog.Errorf("invalid recommendindex")
		}
	}

	if err := recommendmwcli.CreateRecommend(ctx, &recommendmwpb.RecommendReq{
		EntID:          h.EntID,
		RecommenderID:  h.RecommenderID,
		AppGoodID:      h.AppGoodID,
		RecommendIndex: h.RecommendIndex,
		Message:        h.Message,
	}); err != nil {
		return nil, err
	}

	return h.GetRecommend(ctx)
}
