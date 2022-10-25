package recommend

import (
	"context"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/good"
	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/recommend"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/recommend"

	appusermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appusermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.Recommend, error) {
	info, err := mgrcli.CreateRecommend(ctx, &mgrpb.RecommendReq{
		AppID:          &in.AppID,
		GoodID:         &in.GoodID,
		RecommenderID:  &in.RecommenderID,
		Message:        &in.Message,
		RecommendIndex: &in.RecommendIndex,
	})
	if err != nil {
		return nil, err
	}

	return GetRecommend(ctx, info.ID)
}

func UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.Recommend, error) {
	info, err := mgrcli.UpdateRecommend(ctx, &mgrpb.RecommendReq{
		ID:             &in.ID,
		AppID:          &in.AppID,
		Message:        in.Message,
		RecommendIndex: in.RecommendIndex,
	})
	if err != nil {
		return nil, err
	}

	return GetRecommend(ctx, info.ID)
}

func GetRecommends(ctx context.Context, appID string, offset, limit int32) ([]*npool.Recommend, uint32, error) {
	infos, total, err := mgrcli.GetRecommends(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	goodIDs := []string{}
	userIDs := []string{}
	for _, val := range infos {
		goodIDs = append(goodIDs, val.GoodID)
		userIDs = append(userIDs, val.RecommenderID)
	}

	goods, _, err := goodmgrcli.GetGoods(ctx, &goodmgrpb.Conds{
		IDs: &npoolpb.StringSliceVal{
			Op:    cruder.EQ,
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

	users, _, err := appusermwcli.GetManyUsers(ctx, userIDs)
	if err != nil {
		return nil, 0, err
	}

	userMap := map[string]*appusermwpb.User{}
	for _, val := range users {
		userMap[val.ID] = val
	}

	recommends := []*npool.Recommend{}
	for _, info := range infos {
		recommend := &npool.Recommend{
			ID:                      info.ID,
			AppID:                   info.AppID,
			GoodID:                  info.GoodID,
			RecommenderID:           info.RecommenderID,
			RecommenderUsername:     "",
			RecommenderFirstName:    "",
			RecommenderLastName:     "",
			RecommenderEmailAddress: "",
			RecommenderPhoneNo:      "",
			Message:                 "",
			RecommendIndex:          0,
			CreatedAt:               info.CreatedAt,
			UpdatedAt:               info.UpdatedAt,
		}

		good, ok := goodMap[info.GoodID]
		if ok {
			recommend.GoodName = good.Title
		}

		user, ok := userMap[info.RecommenderID]
		if ok {
			recommend.RecommenderUsername = user.Username
			recommend.RecommenderFirstName = user.FirstName
			recommend.RecommenderLastName = user.LastName
			recommend.RecommenderEmailAddress = user.EmailAddress
			recommend.RecommenderPhoneNo = user.PhoneNO
		}

		recommends = append(recommends, recommend)
	}
	return recommends, total, nil
}

func GetRecommend(ctx context.Context, id string) (*npool.Recommend, error) {
	info, err := mgrcli.GetRecommend(ctx, id)
	if err != nil {
		return nil, err
	}

	recommend := &npool.Recommend{
		ID:             info.ID,
		AppID:          info.AppID,
		GoodID:         info.GoodID,
		RecommenderID:  info.RecommenderID,
		Message:        info.Message,
		RecommendIndex: info.RecommendIndex,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}

	good, err := goodmgrcli.GetGood(ctx, info.GoodID)
	if err != nil {
		return nil, err
	}
	if good != nil {
		recommend.GoodName = good.Title
	}

	user, err := appusermwcli.GetUser(ctx, info.AppID, info.RecommenderID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		recommend.RecommenderUsername = user.Username
		recommend.RecommenderFirstName = user.FirstName
		recommend.RecommenderLastName = user.LastName
		recommend.RecommenderEmailAddress = user.EmailAddress
		recommend.RecommenderPhoneNo = user.PhoneNO
	}

	return recommend, nil
}
