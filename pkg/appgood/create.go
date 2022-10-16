package appgood

import (
	"context"
	appusermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/subgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	subgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"
	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/appgood"
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

	coinType, err := coininfocli.GetCoinInfo(ctx, info.CoinTypeID)
	if err != nil {
		return nil, err
	}

	supportCoins := []*npool.Good_CoinInfo{}
	for _, val := range info.SupportCoinTypeIDs {
		coinTypeInfo, err := coininfocli.GetCoinInfo(ctx, val)
		if err != nil {
			return nil, err
		}
		supportCoins = append(supportCoins, &npool.Good_CoinInfo{
			CoinTypeID:  info.CoinTypeID,
			CoinLogo:    coinTypeInfo.Logo,
			CoinName:    coinTypeInfo.Name,
			CoinUnit:    coinTypeInfo.Unit,
			CoinPreSale: coinTypeInfo.PreSale,
		})
	}

	var recommenderEmailAddress *string
	var recommenderPhoneNO *string
	var recommenderUsername *string
	var recommenderFirstName *string
	var recommenderLastName *string

	if info.RecommenderID != nil {
		recommender, err := appusermwcli.GetUser(ctx, info.AppID, info.GetRecommenderID())
		if err != nil {
			return nil, err
		}

		recommenderEmailAddress = &recommender.EmailAddress
		recommenderPhoneNO = &recommender.PhoneNO
		recommenderUsername = &recommender.Username
		recommenderFirstName = &recommender.FirstName
		recommenderLastName = &recommender.LastName
	}

	subGoods, _, err := goodmgrcli.GetSubGoods(ctx, &subgoodmgrpb.Conds{
		MainGoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GoodID,
		},
	}, 0, 0)

	subGoodIDs := []string{}
	for _, val := range subGoods {
		subGoodIDs = append(subGoodIDs, val.SubGoodID)
	}

	//subappGoods, _, err := appgoodmwcli.GetGoods(ctx, &appgoodmgrpb.Conds{
	//	ID:     nil,
	//	AppID:  nil,
	//	GoodID: nil,
	//}, 0, int32(len(subGoodIDs)))
	//if err != nil {
	//	return nil, err
	//}

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
		CoinLogo:                coinType.Logo,
		CoinName:                coinType.Name,
		CoinUnit:                coinType.Unit,
		CoinPreSale:             coinType.PreSale,
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
		GoodTotal:               info.GoodTotal,
		GoodLocked:              info.GoodLocked,
		GoodInService:           info.GoodInService,
		GoodSold:                info.GoodSold,
		SubGoods:                nil,
		Must:                    true,
		Commission:              true,
		StartAt:                 info.StartAt,
		CreatedAt:               info.CreatedAt,
	}

	// TODO: expand recommender
	// TODO: expand coin type
	// TODO: expand support coin types

	return info1, nil
}
