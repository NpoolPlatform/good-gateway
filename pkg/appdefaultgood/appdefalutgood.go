//nolint
package appdefaultgood

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appdefaultgood"
	appdefaultgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	appdefaultgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appdefaultgood"

	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
)

func CreateAppDefaultGood(
	ctx context.Context,
	appID, goodID, coinTypeID string,
) (*npool.AppDefaultGood, error) {
	info, err := appdefaultgoodmgrcli.CreateAppDefaultGood(ctx, &appdefaultgoodmgrpb.AppDefaultGoodReq{
		AppID:      &appID,
		GoodID:     &goodID,
		CoinTypeID: &coinTypeID,
	})
	if err != nil {
		return nil, err
	}
	return GetAppDefaultGood(ctx, info.ID)
}

func UpdateAppDefaultGood(
	ctx context.Context,
	id, goodID string,
) (*npool.AppDefaultGood, error) {
	info, err := appdefaultgoodmgrcli.UpdateAppDefaultGood(ctx, &appdefaultgoodmgrpb.AppDefaultGoodReq{
		ID:     &id,
		GoodID: &goodID,
	})
	if err != nil {
		return nil, err
	}
	return GetAppDefaultGood(ctx, info.ID)
}

func DeleteAppDefaultGood(
	ctx context.Context,
	id string,
) (*npool.AppDefaultGood, error) {
	info, err := appdefaultgoodmgrcli.DeleteAppDefaultGood(ctx, id)
	if err != nil {
		return nil, err
	}
	return &npool.AppDefaultGood{
		ID:         info.ID,
		AppID:      info.AppID,
		GoodID:     info.GoodID,
		CoinTypeID: info.CoinTypeID,
		CreatedAt:  info.CreatedAt,
		UpdatedAt:  info.UpdatedAt,
	}, nil
}

func GetAppDefaultGood(ctx context.Context, id string) (*npool.AppDefaultGood, error) {
	info, err := appdefaultgoodmgrcli.GetAppDefaultGood(ctx, id)
	if err != nil {
		return nil, err
	}
	coinInfo, err := coincli.GetCoin(ctx, info.CoinTypeID)
	if err != nil {
		return nil, err
	}
	return &npool.AppDefaultGood{
		ID:         info.ID,
		AppID:      info.AppID,
		GoodID:     info.GoodID,
		CoinTypeID: info.CoinTypeID,
		CoinUnit:   coinInfo.Unit,
		CreatedAt:  info.CreatedAt,
		UpdatedAt:  info.UpdatedAt,
	}, nil
}

func GetAppDefaultGoods(ctx context.Context, appID string, offset, limit int32) ([]*npool.AppDefaultGood, uint32, error) {
	rows, total, err := appdefaultgoodmgrcli.GetAppDefaultGoods(ctx, &appdefaultgoodmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	coinTypeIDs := []string{}
	for _, id := range rows {
		coinTypeIDs = append(coinTypeIDs, id.CoinTypeID)
	}
	coinInfos, _, err := coincli.GetCoins(ctx, &coinpb.Conds{
		IDs: &basetypes.StringSliceVal{
			Op:    cruder.IN,
			Value: coinTypeIDs,
		},
	}, 0, int32(len(coinTypeIDs)))
	if err != nil {
		return nil, 0, err
	}
	coinMap := map[string]*coinpb.Coin{}
	for _, val := range coinInfos {
		coinMap[val.ID] = val
	}
	var infos []*npool.AppDefaultGood
	{
	}
	for _, val := range rows {
		coinInfo, ok := coinMap[val.CoinTypeID]
		if !ok {
			continue
		}
		infos = append(infos, &npool.AppDefaultGood{
			ID:         val.ID,
			AppID:      val.AppID,
			GoodID:     val.GoodID,
			CoinTypeID: val.CoinTypeID,
			CoinUnit:   coinInfo.Unit,
			CreatedAt:  val.CreatedAt,
			UpdatedAt:  val.UpdatedAt,
		})
	}
	return infos, total, nil
}
