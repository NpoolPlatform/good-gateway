package recommend

import (
	"context"
	"fmt"

	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/recommend"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"
)

func (h *Handler) DeleteRecommend(ctx context.Context) (*npool.Recommend, error) {
	exist, err := recommendmwcli.ExistRecommendConds(ctx, &recommendmwpb.Conds{
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: *h.ID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid recommend")
	}

	info, err := h.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := recommendmwcli.DeleteRecommend(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
