package promotion

import (
	"context"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/good"
	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/promotion"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/promotion"
)

func CreatePromotion(ctx context.Context, in *npool.CreatePromotionRequest) (*npool.Promotion, error) {
	info, err := mgrcli.CreatePromotion(ctx, &mgrpb.PromotionReq{
		AppID:   &in.AppID,
		GoodID:  &in.GoodID,
		Message: &in.Message,
		StartAt: &in.StartAt,
		EndAt:   &in.EndAt,
		Price:   &in.Price,
		Posters: in.Posters,
	})
	if err != nil {
		return nil, err
	}

	return GetPromotion(ctx, info.ID)
}

func UpdatePromotion(ctx context.Context, in *npool.UpdatePromotionRequest) (*npool.Promotion, error) {
	info, err := mgrcli.UpdatePromotion(ctx, &mgrpb.PromotionReq{
		ID:      &in.ID,
		AppID:   &in.AppID,
		Message: in.Message,
		StartAt: in.StartAt,
		EndAt:   in.EndAt,
		Price:   in.Price,
		Posters: in.Posters,
	})
	if err != nil {
		return nil, err
	}

	return GetPromotion(ctx, info.ID)
}

func GetPromotions(ctx context.Context, appID string, offset, limit int32) ([]*npool.Promotion, uint32, error) {
	infos, total, err := mgrcli.GetPromotions(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return nil, 0, err
	}

	goodIDs := []string{}
	for _, val := range infos {
		goodIDs = append(goodIDs, val.GoodID)
	}

	goods, _, err := goodmgrcli.GetGoods(ctx, &goodmgrpb.Conds{
		IDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: goodIDs,
		},
	}, 0, int32(len(goodIDs)))
	if err != nil {
		return nil, 0, err
	}

	goodMap := map[string]*goodmgrpb.Good{}
	for _, val := range goods {
		goodMap[val.ID] = val
	}

	promotions := []*npool.Promotion{}
	for _, info := range infos {
		promotion := &npool.Promotion{
			ID:        info.ID,
			AppID:     info.AppID,
			GoodID:    info.GoodID,
			Message:   info.Message,
			StartAt:   info.StartAt,
			EndAt:     info.EndAt,
			Price:     info.Price,
			Posters:   info.Posters,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		}

		good, ok := goodMap[info.GoodID]
		if ok {
			promotion.GoodName = good.Title
		}
		promotions = append(promotions, promotion)
	}
	return promotions, total, nil
}

func GetPromotion(ctx context.Context, id string) (*npool.Promotion, error) {
	info, err := mgrcli.GetPromotion(ctx, id)
	if err != nil {
		return nil, err
	}

	good, err := goodmgrcli.GetGood(ctx, info.GoodID)
	if err != nil {
		return nil, err
	}

	return &npool.Promotion{
		ID:        info.ID,
		AppID:     info.AppID,
		GoodID:    info.GoodID,
		GoodName:  good.Title,
		Message:   info.Message,
		StartAt:   info.StartAt,
		EndAt:     info.EndAt,
		Price:     info.Price,
		Posters:   info.Posters,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}, nil
}
