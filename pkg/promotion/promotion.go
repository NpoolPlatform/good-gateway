package promotion

import (
	"context"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/promotion"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/promotion"

	appgoodpb "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"

	appgoodm "github.com/NpoolPlatform/good-gateway/pkg/appgood"
)

func CreatePromotion(ctx context.Context, in *npool.CreatePromotionRequest) (*appgoodpb.Good, error) {
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

	goodmw, err := appgoodmwcli.GetGood(ctx, info.GoodID)
	if err != nil {
		return nil, err
	}

	good, err := appgoodm.Scan(ctx, goodmw)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func UpdatePromotion(ctx context.Context, in *npool.UpdatePromotionRequest) (*appgoodpb.Good, error) {
	info, err := mgrcli.UpdatePromotion(ctx, &mgrpb.PromotionReq{
		ID:      &in.ID,
		Message: in.Message,
		StartAt: in.StartAt,
		EndAt:   in.EndAt,
		Price:   in.Price,
		Posters: in.Posters,
	})
	if err != nil {
		return nil, err
	}

	goodmw, err := appgoodmwcli.GetGood(ctx, info.GoodID)
	if err != nil {
		return nil, err
	}

	good, err := appgoodm.Scan(ctx, goodmw)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func GetPromotions(ctx context.Context, appID string, offset, limit int32) ([]*appgoodpb.Good, uint32, error) {
	infos, total, err := mgrcli.GetPromotions(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	goodIDs := []string{}
	for _, val := range infos {
		goodIDs = append(goodIDs, val.GoodID)
	}

	goodmw, _, err := appgoodmwcli.GetGoods(ctx, &appgoodmgrpb.Conds{
		GoodIDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: goodIDs,
		},
	}, 0, int32(len(infos)))
	if err != nil {
		return nil, 0, err
	}

	good, err := appgoodm.Scans(ctx, goodmw, appID)
	if err != nil {
		return nil, 0, err
	}

	return good, total, nil
}
