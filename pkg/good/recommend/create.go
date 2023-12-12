package recommend

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"

	"github.com/google/uuid"
)

func (h *Handler) CreateRecommend(ctx context.Context) (*npool.Recommend, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.RecommenderID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := recommendmwcli.CreateRecommend(ctx, &recommendmwpb.RecommendReq{
		EntID:          h.EntID,
		AppID:          h.AppID,
		RecommenderID:  h.RecommenderID,
		GoodID:         h.GoodID,
		RecommendIndex: h.RecommendIndex,
		Message:        h.Message,
	}); err != nil {
		return nil, err
	}

	return h.GetRecommend(ctx)
}
