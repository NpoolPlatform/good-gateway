package subgood

import (
	"context"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/subgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/subgood"

	npoolpb "github.com/NpoolPlatform/message/npool"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/good"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
)

func CreateSubGood(ctx context.Context, in *npool.CreateSubGoodRequest) (*npool.SubGood, error) {
	info, err := mgrcli.CreateSubGood(ctx, &mgrpb.SubGoodReq{
		AppID:      &in.AppID,
		MainGoodID: &in.MainGoodID,
		SubGoodID:  &in.SubGoodID,
		Must:       &in.Must,
		Commission: &in.Commission,
	})
	if err != nil {
		return nil, err
	}

	return GetSubGood(ctx, info.ID)
}

func UpdateSubGood(ctx context.Context, in *npool.UpdateSubGoodRequest) (*npool.SubGood, error) {
	info, err := mgrcli.UpdateSubGood(ctx, &mgrpb.SubGoodReq{
		ID:         &in.ID,
		SubGoodID:  in.SubGoodID,
		Must:       in.Must,
		Commission: in.Commission,
	})
	if err != nil {
		return nil, err
	}

	return GetSubGood(ctx, info.ID)
}

func GetSubGood(ctx context.Context, id string) (*npool.SubGood, error) {
	info, err := mgrcli.GetSubGood(ctx, id)
	if err != nil {
		return nil, err
	}

	goodIDs := []string{info.MainGoodID, info.SubGoodID}
	goods, _, err := goodmgrcli.GetGoods(ctx, &goodmgrpb.Conds{
		IDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: goodIDs,
		},
	}, 0, int32(len(goodIDs)))
	if err != nil {
		return nil, err
	}

	mainGoodName := ""
	subGoodName := ""
	for _, val := range goods {
		if val.ID == info.MainGoodID {
			mainGoodName = val.Title
		}
		if val.ID == info.SubGoodID {
			subGoodName = val.Title
		}
	}

	return &npool.SubGood{
		ID:           info.ID,
		AppID:        info.AppID,
		MainGoodID:   info.MainGoodID,
		MainGoodName: mainGoodName,
		SubGoodID:    info.SubGoodID,
		SubGoodName:  subGoodName,
		Must:         info.Must,
		Commission:   info.Commission,
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
	}, nil
}

func GetSubGoods(ctx context.Context, appID string, offset, limit int32) ([]*npool.SubGood, uint32, error) {
	infos, total, err := mgrcli.GetSubGoods(ctx, &mgrpb.Conds{
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
		goodIDs = append(goodIDs, val.SubGoodID, val.MainGoodID)
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

	subGoods := []*npool.SubGood{}
	for _, info := range infos {
		mainGoodName := ""
		subGoodName := ""
		for _, val := range goods {
			if val.ID == info.MainGoodID {
				mainGoodName = val.Title
			}
			if val.ID == info.SubGoodID {
				subGoodName = val.Title
			}
		}
		subGoods = append(subGoods, &npool.SubGood{
			ID:           info.ID,
			AppID:        info.AppID,
			MainGoodID:   info.MainGoodID,
			MainGoodName: mainGoodName,
			SubGoodID:    info.SubGoodID,
			SubGoodName:  subGoodName,
			Must:         info.Must,
			Commission:   info.Commission,
			CreatedAt:    info.CreatedAt,
			UpdatedAt:    info.UpdatedAt,
		})
	}
	return subGoods, total, nil
}
