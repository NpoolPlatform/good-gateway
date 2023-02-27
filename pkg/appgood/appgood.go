//nolint
package appgood

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/subgood"
	subgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	ordermgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/order"

	appcoininfocli "github.com/NpoolPlatform/chain-middleware/pkg/client/appcoin"
	appcoinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/appgood"

	appusermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appusermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
)

func CreateAppGood(
	ctx context.Context,
	appID, goodID string, online, visible bool,
	goodName string, price string,
	displayIndex, purchaseLimit, commissionPercent int32,
	saleStart, saleEnd, serviceStart *uint32,
	techFeeRatio, elecFeeRatio *uint32,
	commSettleType commmgrpb.SettleType,
	openPurchase, intoProductPage bool,
	cancelableBefore uint32,
	userPurchaseLimit string,
) (*npool.Good, error) {
	info, err := appgoodmwcli.CreateGood(ctx, &appgoodmgrpb.AppGoodReq{
		AppID:                &appID,
		GoodID:               &goodID,
		Online:               &online,
		Visible:              &visible,
		GoodName:             &goodName,
		Price:                &price,
		DisplayIndex:         &displayIndex,
		PurchaseLimit:        &purchaseLimit,
		CommissionPercent:    &commissionPercent,
		SaleStartAt:          saleStart,
		SaleEndAt:            saleEnd,
		ServiceStartAt:       serviceStart,
		TechnicalFeeRatio:    techFeeRatio,
		ElectricityFeeRatio:  elecFeeRatio,
		CommissionSettleType: &commSettleType,
		OpenPurchase:         &openPurchase,
		IntoProductPage:      &intoProductPage,
		CancelableBefore:     &cancelableBefore,
		UserPurchaseLimit:    &userPurchaseLimit,
	})
	if err != nil {
		return nil, err
	}

	return Scan(ctx, info)
}

func GetAppGoods(ctx context.Context, appID string, offset, limit int32) ([]*npool.Good, uint32, error) {
	goods, total, err := appgoodmwcli.GetGoods(ctx, &appgoodmgrpb.Conds{
		AppID: &commonpb.StringVal{
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
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		GoodID: &commonpb.StringVal{
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
		ID:                   &in.ID,
		Online:               in.Online,
		Visible:              in.Visible,
		GoodName:             in.GoodName,
		Price:                in.Price,
		DisplayIndex:         in.DisplayIndex,
		PurchaseLimit:        in.PurchaseLimit,
		CommissionPercent:    in.CommissionPercent,
		SaleStartAt:          in.SaleStartAt,
		SaleEndAt:            in.SaleEndAt,
		ServiceStartAt:       in.ServiceStartAt,
		TechnicalFeeRatio:    in.TechnicalFeeRatio,
		ElectricityFeeRatio:  in.ElectricityFeeRatio,
		DailyRewardAmount:    in.DailyRewardAmount,
		CommissionSettleType: in.CommissionSettleType,
		Descriptions:         in.Descriptions,
		GoodBanner:           in.GoodBanner,
		DisplayNames:         in.DisplayNames,
		OpenPurchase:         in.OpenPurchase,
		IntoProductPage:      in.IntoProductPage,
		CancelableBefore:     in.CancelableBefore,
		UserPurchaseLimit:    in.UserPurchaseLimit,
	})
	if err != nil {
		return nil, err
	}

	if in.ServiceStartAt != nil {
		offset := int32(0)
		limit := constant.DefaultRowLimit

		for {
			orders, _, err := ordermwcli.GetOrders(ctx, &ordermwpb.Conds{
				AppID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: info.AppID,
				},
				GoodID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: info.GoodID,
				},
			}, offset, limit)
			if err != nil {
				return nil, err
			}
			if len(orders) == 0 {
				break
			}

			reqs := []*ordermwpb.OrderReq{}
			for _, ord := range orders {
				if ord.Start > in.GetServiceStartAt() {
					continue
				}
				if ord.PaymentID == uuid1.InvalidUUIDStr || ord.PaymentID == "" {
					switch ord.OrderState {
					case ordermgrpb.OrderState_Paid:
						fallthrough //nolint
					case ordermgrpb.OrderState_InService:
						fallthrough //nolint
					case ordermgrpb.OrderState_Expired:
						logger.Sugar().Warnw("UpdateAppGood", "OrderID", ord.ID, "State", ord.OrderState, "PaymentID", ord.PaymentID)
					}
					continue
				}
				reqs = append(reqs, &ordermwpb.OrderReq{
					ID:        &ord.ID,
					PaymentID: &ord.PaymentID,
					Start:     in.ServiceStartAt,
				})
			}

			if len(reqs) > 0 {
				_, err = ordermwcli.UpdateOrders(ctx, reqs)
				if err != nil {
					return nil, err
				}
			}

			offset += limit
		}
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
	coinTypeIDs := []string{}
	for _, val := range infos {
		coinTypeIDs = append(coinTypeIDs, val.CoinTypeID)
	}

	coins, _, err := appcoininfocli.GetCoins(ctx, &appcoinpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		CoinTypeIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: coinTypeIDs,
		},
	}, 0, int32(len(coinTypeIDs)))
	if err != nil {
		return nil, err
	}

	ctMap := map[string]*appcoinpb.Coin{}
	for _, coin := range coins {
		ctMap[coin.CoinTypeID] = coin
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
	ctMap map[string]*appcoinpb.Coin,
	userMap map[string]*appusermwpb.User,
	goodIDs []string,
	appID string,
) (map[string][]*npool.Good, error) {
	subGoods, _, err := goodmgrcli.GetSubGoods(ctx, &subgoodmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		MainGoodIDs: &commonpb.StringSliceVal{
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
			AppID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			GoodIDs: &commonpb.StringSliceVal{
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
	ctMap map[string]*appcoinpb.Coin,
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
					CoinPreSale:  coinTypeInfo.Presale,
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
			GoodType:                info.GoodType,
			BenefitType:             info.BenefitType,
			GoodName:                info.GoodName,
			Unit:                    info.Unit,
			UnitAmount:              info.UnitAmount,
			BenefitIntervalHours:    info.BenefitIntervalHours,
			TestOnly:                info.TestOnly,
			Posters:                 info.Posters,
			Labels:                  info.Labels,
			VoteCount:               info.VoteCount,
			Rating:                  info.Rating,
			SupportCoins:            supportCoins,
			Total:                   info.GoodTotal,
			Locked:                  info.GoodLocked,
			InService:               info.GoodInService,
			WaitStart:               info.GoodWaitStart,
			Sold:                    info.GoodSold,
			StartAt:                 info.StartAt,
			CreatedAt:               info.CreatedAt,
			SaleStartAt:             info.SaleStartAt,
			SaleEndAt:               info.SaleEndAt,
			ServiceStartAt:          info.ServiceStartAt,
			TechnicalFeeRatio:       info.TechnicalFeeRatio,
			ElectricityFeeRatio:     info.ElectricityFeeRatio,
			DailyRewardAmount:       info.DailyRewardAmount,
			CommissionSettleType:    info.CommissionSettleType,
			Descriptions:            info.Descriptions,
			GoodBanner:              info.GoodBanner,
			DisplayNames:            info.DisplayNames,
			OpenPurchase:            info.OpenPurchase,
			IntoProductPage:         info.IntoProductPage,
			CancelableBefore:        info.CancelableBefore,
			UserPurchaseLimit:       info.UserPurchaseLimit,
		}

		coinType, ok := ctMap[info.CoinTypeID]
		if ok {
			info1.CoinLogo = coinType.Logo
			info1.CoinName = coinType.Name
			info1.CoinUnit = coinType.Unit
			info1.CoinPreSale = coinType.Presale
			info1.CoinEnv = coinType.ENV
			info1.CoinHomePage = coinType.HomePage
			info1.CoinSpecs = coinType.Specs
		}

		goods = append(goods, info1)
	}
	return goods
}
