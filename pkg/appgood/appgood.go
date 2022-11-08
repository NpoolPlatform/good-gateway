//nolint
package appgood

import (
	"context"
	"fmt"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/subgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	subgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/appgood"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"

	appusermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appusermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func CreateAppGood(
	ctx context.Context,
	appID, goodID string, online, visible bool,
	goodName string, price string,
	displayIndex, purchaseLimit, commissionPercent int32,
) (*npool.Good, error) {
	info, err := appgoodmwcli.CreateGood(ctx, &appgoodmgrpb.AppGoodReq{
		AppID:             &appID,
		GoodID:            &goodID,
		Online:            &online,
		Visible:           &visible,
		GoodName:          &goodName,
		Price:             &price,
		DisplayIndex:      &displayIndex,
		PurchaseLimit:     &purchaseLimit,
		CommissionPercent: &commissionPercent,
	})
	if err != nil {
		return nil, err
	}

	return Scan(ctx, info)
}

func GetAppGoods(ctx context.Context, appID string, offset, limit int32) ([]*npool.Good, uint32, error) {
	goods, total, err := appgoodmwcli.GetGoods(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	infos, err := Scans(ctx, goods, appID)
	if err != nil {
		return nil, 0, err
	}

	return infos, total, err
}

func GetAppGood(ctx context.Context, appID, goodID string) (*npool.Good, error) {
	good, err := appgoodmwcli.GetGoodOnly(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: goodID,
		},
	})
	if err != nil {
		return nil, err
	}

	if good == nil {
		return nil, fmt.Errorf("app good is invalid")
	}

	info, err := Scan(ctx, good)
	if err != nil {
		return nil, err
	}

	return info, err
}

func UpdateAppGood(ctx context.Context, in *npool.UpdateAppGoodRequest) (*npool.Good, error) {
	info, err := appgoodmwcli.UpdateGood(ctx, &appgoodmgrpb.AppGoodReq{
		ID:                &in.ID,
		Online:            in.Online,
		Visible:           in.Visible,
		GoodName:          in.GoodName,
		Price:             in.Price,
		DisplayIndex:      in.DisplayIndex,
		PurchaseLimit:     in.PurchaseLimit,
		CommissionPercent: in.CommissionPercent,
	})
	if err != nil {
		return nil, err
	}

	return Scan(ctx, info)
}

func Scan(ctx context.Context, info *goodmwpb.Good) (*npool.Good, error) {
	infos := []*goodmwpb.Good{}
	infos = append(infos, info)
	goods, err := Scans(ctx, infos, info.AppID)
	if err != nil {
		return nil, err
	}
	if len(goods) == 0 {
		return nil, fmt.Errorf("goods invalid")
	}
	return goods[0], nil
}

func Scans(ctx context.Context, infos []*goodmwpb.Good, appID string) ([]*npool.Good, error) {
	coinTypes, err := coininfocli.GetCoinInfos(ctx, nil)
	if err != nil {
		return nil, err
	}

	ctMap := map[string]*coininfopb.CoinInfo{}
	for _, coinType := range coinTypes {
		ctMap[coinType.ID] = coinType
	}

	userIDs := []string{}
	goodIDs := []string{}
	for _, val := range infos {
		if val.RecommenderID != nil {
			userIDs = append(userIDs, *val.RecommenderID)
		}
		goodIDs = append(goodIDs, val.GoodID)
	}

	userInfos := []*appusermwpb.User{}

	if len(userIDs) > 0 {
		userInfos, _, err = appusermwcli.GetManyUsers(ctx, userIDs)
		if err != nil {
			return nil, err
		}
	}

	userMap := map[string]*appusermwpb.User{}
	for _, userInfo := range userInfos {
		userMap[userInfo.ID] = userInfo
	}

	subGoodMap := map[string][]*npool.Good{}
	if len(goodIDs) > 0 {
		subGoodMap, err = getSubGoods(ctx, ctMap, userMap, goodIDs, appID)
		if err != nil {
			return nil, err
		}
	}

	goods := getGoodInfos(ctMap, userMap, infos)

	for key := range goods {
		subGood, ok := subGoodMap[goods[key].GoodID]
		if ok {
			goods[key].SubGoods = subGood
		}
	}
	return goods, nil
}

func getSubGoods(
	ctx context.Context,
	ctMap map[string]*coininfopb.CoinInfo,
	userMap map[string]*appusermwpb.User,
	goodIDs []string,
	appID string,
) (map[string][]*npool.Good, error) {
	subGoods, _, err := goodmgrcli.GetSubGoods(ctx, &subgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		MainGoodIDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: goodIDs,
		},
	}, 0, int32(len(goodIDs)))

	if err != nil {
		return nil, err
	}

	subGoodIDs := []string{}
	for _, val := range subGoods {
		subGoodIDs = append(subGoodIDs, val.SubGoodID)
	}

	subAppGoods := []*goodmwpb.Good{}
	if len(subGoodIDs) > 0 {
		subAppGoods, _, err = appgoodmwcli.GetGoods(ctx, &appgoodmgrpb.Conds{
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			GoodIDs: &npoolpb.StringSliceVal{
				Op:    cruder.IN,
				Value: subGoodIDs,
			},
		}, 0, int32(len(subGoodIDs)))
		if err != nil {
			return nil, err
		}
	}

	goods := getGoodInfos(ctMap, userMap, subAppGoods)
	subGoodMap := map[string][]*npool.Good{}
	for _, subGood := range subGoods {
		subGoodInfos := []*npool.Good{}
		for _, subGood1 := range subGoods {
			for _, good := range goods {
				if subGood.MainGoodID == subGood1.MainGoodID && subGood1.SubGoodID == good.GoodID {
					good.Commission = subGood1.Commission
					good.Must = subGood1.Must
					subGoodInfos = append(subGoodInfos, good)
				}
			}
		}
		subGoodMap[subGood.MainGoodID] = subGoodInfos
	}
	return subGoodMap, nil
}

func getGoodInfos(
	ctMap map[string]*coininfopb.CoinInfo,
	userMap map[string]*appusermwpb.User,
	infos []*goodmwpb.Good,
) []*npool.Good {

	goods := []*npool.Good{}

	for _, info := range infos {
		supportCoins := []*npool.Good_CoinInfo{}
		for _, supportCoinTypeID := range info.SupportCoinTypeIDs {
			coinTypeInfo, ok := ctMap[supportCoinTypeID]
			if ok {
				supportCoins = append(supportCoins, &npool.Good_CoinInfo{
					CoinTypeID:   info.CoinTypeID,
					CoinLogo:     coinTypeInfo.Logo,
					CoinName:     coinTypeInfo.Name,
					CoinUnit:     coinTypeInfo.Unit,
					CoinPreSale:  coinTypeInfo.PreSale,
					CoinEnv:      coinTypeInfo.ENV,
					CoinHomePage: coinTypeInfo.HomePage,
					CoinSpecs:    coinTypeInfo.Specs,
				})
			}
		}

		var recommenderEmailAddress *string
		var recommenderPhoneNO *string
		var recommenderUsername *string
		var recommenderFirstName *string
		var recommenderLastName *string

		if info.RecommenderID != nil {
			userInfo, ok := userMap[*info.RecommenderID]
			if ok {
				recommenderEmailAddress = &userInfo.EmailAddress
				recommenderPhoneNO = &userInfo.PhoneNO
				recommenderUsername = &userInfo.Username
				recommenderFirstName = &userInfo.FirstName
				recommenderLastName = &userInfo.LastName
			}
		}

		info1 := &npool.Good{
			ID:                      info.ID,
			AppID:                   info.AppID,
			GoodID:                  info.GoodID,
			Online:                  info.Online,
			Visible:                 info.Visible,
			Price:                   info.Price,
			DisplayIndex:            info.DisplayIndex,
			PurchaseLimit:           info.PurchaseLimit,
			CommissionPercent:       info.CommissionPercent,
			PromotionStartAt:        info.PromotionStartAt,
			PromotionEndAt:          info.PromotionEndAt,
			PromotionMessage:        info.PromotionMessage,
			PromotionPrice:          info.PromotionPrice,
			PromotionPosters:        info.PromotionPosters,
			RecommenderID:           info.RecommenderID,
			RecommenderEmailAddress: recommenderEmailAddress,
			RecommenderPhoneNO:      recommenderPhoneNO,
			RecommenderUsername:     recommenderUsername,
			RecommenderFirstName:    recommenderFirstName,
			RecommenderLastName:     recommenderLastName,
			RecommendMessage:        info.RecommendMessage,
			RecommendIndex:          info.RecommendIndex,
			RecommendAt:             info.RecommendAt,
			DeviceType:              info.DeviceType,
			DeviceManufacturer:      info.DeviceManufacturer,
			DevicePowerComsuption:   info.DevicePowerComsuption,
			DeviceShipmentAt:        info.DeviceShipmentAt,
			DevicePosters:           info.DevicePosters,
			DurationDays:            info.DurationDays,
			VendorLocationCountry:   info.VendorLocationCountry,
			CoinTypeID:              info.CoinTypeID,
			GoodType:                info.GoodType,
			BenefitType:             info.BenefitType,
			GoodName:                info.GoodName,
			Unit:                    info.Unit,
			UnitAmount:              info.UnitAmount,
			TestOnly:                info.TestOnly,
			Posters:                 info.Posters,
			Labels:                  info.Labels,
			VoteCount:               info.VoteCount,
			Rating:                  info.Rating,
			SupportCoins:            supportCoins,
			Total:                   info.GoodTotal,
			Locked:                  info.GoodLocked,
			InService:               info.GoodInService,
			Sold:                    info.GoodSold,
			SubGoods:                nil,
			Must:                    true,
			Commission:              true,
			StartAt:                 info.StartAt,
			CreatedAt:               info.CreatedAt,
		}

		coinType, ok := ctMap[info.CoinTypeID]
		if ok {
			info1.CoinLogo = coinType.Logo
			info1.CoinName = coinType.Name
			info1.CoinUnit = coinType.Unit
			info1.CoinPreSale = coinType.PreSale
			info1.CoinEnv = coinType.ENV
			info1.CoinHomePage = coinType.HomePage
			info1.CoinSpecs = coinType.Specs
		}

		goods = append(goods, info1)
	}
	return goods
}
