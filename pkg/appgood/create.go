package appgood

import (
	"context"

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

	info1 := &npool.Good{
		ID:                    info.ID,
		AppID:                 info.AppID,
		GoodID:                info.GoodID,
		Online:                info.Online,
		Visible:               info.Visible,
		Price:                 info.Price,
		DisplayIndex:          info.DisplayIndex,
		PurchaseLimit:         info.PurchaseLimit,
		CommissionPercent:     info.CommissionPercent,
		PromotionStartAt:      info.PromotionStartAt,
		PromotionEndAt:        info.PromotionEndAt,
		PromotionMessage:      info.PromotionMessage,
		PromotionPrice:        info.PromotionPrice,
		PromotionPosters:      info.PromotionPosters,
		RecommenderID:         info.RecommenderID,
		RecommendMessage:      info.RecommendMessage,
		RecommendIndex:        info.RecommendIndex,
		RecommendAt:           info.RecommendAt,
		DeviceType:            info.DeviceType,
		DeviceManufacturer:    info.DeviceManufacturer,
		DevicePowerComsuption: info.DevicePowerComsuption,
		DeviceShipmentAt:      info.DeviceShipmentAt,
		DevicePosters:         info.DevicePosters,
		DurationDays:          info.DurationDays,
		VendorLocationCountry: info.VendorLocationCountry,
		CoinTypeID:            info.CoinTypeID,
		GoodType:              info.GoodType,
		BenefitType:           info.BenefitType,
		GoodName:              info.GoodName,
		Unit:                  info.Unit,
		UnitAmount:            info.UnitAmount,
		TestOnly:              info.TestOnly,
		Posters:               info.Posters,
		Labels:                info.Labels,
		VoteCount:             info.VoteCount,
		Rating:                info.Rating,
		GoodTotal:             info.GoodTotal,
		GoodLocked:            info.GoodLocked,
		GoodInService:         info.GoodInService,
		GoodSold:              info.GoodSold,
		Must:                  true,
		Commission:            true,
		StartAt:               info.StartAt,
		CreatedAt:             info.CreatedAt,
	}

	return info1, nil
}
