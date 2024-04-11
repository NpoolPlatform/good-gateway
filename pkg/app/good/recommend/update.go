package recommend

import (
	"context"
	"fmt"

	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (h *Handler) UpdateRecommend(ctx context.Context) (*npool.Recommend, error) {
	exist, err := recommendmwcli.ExistRecommendConds(ctx, &recommendmwpb.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid recommend")
	}

	if _, err := recommendmwcli.UpdateRecommend(ctx, &recommendmwpb.RecommendReq{
		ID:      h.ID,
		Message: h.Message,
	}); err != nil {
		return nil, err
	}

	return h.GetRecommend(ctx)
}
