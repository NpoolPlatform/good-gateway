package common

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"

	"github.com/google/uuid"
)

func GetPoolGoodUsers(ctx context.Context, poolGoodUserIDs []string) (map[string]*goodusermwpb.GoodUser, error) {
	for _, poolGoodUserID := range poolGoodUserIDs {
		if _, err := uuid.Parse(poolGoodUserID); err != nil {
			return nil, err
		}
	}

	orderUsers, _, err := goodusermwcli.GetGoodUsers(ctx, &goodusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: poolGoodUserIDs},
	}, int32(0), int32(len(poolGoodUserIDs)))
	if err != nil {
		return nil, err
	}
	orderUserMap := map[string]*goodusermwpb.GoodUser{}
	for _, orderUser := range orderUsers {
		orderUserMap[orderUser.EntID] = orderUser
	}
	return orderUserMap, nil
}
