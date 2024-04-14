package recommend

import (
	"context"

	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/google/uuid"
)

func (h *Handler) CreateRecommend(ctx context.Context) (*npool.Recommend, error) {
	if err := h.CheckUserWithUserID(ctx, *h.RecommenderID); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
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
