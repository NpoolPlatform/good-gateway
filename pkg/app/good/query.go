package good

import (
	"context"
	"fmt"

	appcoinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/app/coin"
	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	goods    []*appgoodmwpb.Good
	infos    []*npool.Good
	appCoins map[string]*appcoinmwpb.Coin
}

func (h *queryHandler) getCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, good := range h.goods {
		if _, err := uuid.Parse(good.CoinTypeID); err != nil {
			continue
		}
		coinTypeIDs = append(coinTypeIDs, good.CoinTypeID)
	}
	coins, _, err := appcoinmwcli.GetCoins(ctx, &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}, int32(0), int32(len(coinTypeIDs)))
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.appCoins[coin.CoinTypeID] = coin
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, good := range h.goods {
		info := &npool.Good{
			ID:                     good.ID,
			EntID:                  good.EntID,
			AppID:                  good.AppID,
			GoodID:                 good.GoodID,
			Online:                 good.Online,
			Visible:                good.Visible,
			UnitPrice:              good.UnitPrice,
			PackagePrice:           good.PackagePrice,
			DisplayIndex:           good.DisplayIndex,
			DeviceType:             good.DeviceType,
			DeviceManufacturer:     good.DeviceManufacturer,
			DevicePowerConsumption: good.DevicePowerConsumption,
			DeviceShipmentAt:       good.DeviceShipmentAt,
			DevicePosters:          good.DevicePosters,
			CoinTypeID:             good.CoinTypeID,
			VendorLocationCountry:  good.VendorLocationCountry,
			VendorBrandName:        good.VendorBrandName,
			VendorBrandLogo:        good.VendorBrandLogo,
			GoodType:               good.GoodType,
			BenefitType:            good.BenefitType,
			GoodName:               good.GoodName,
			QuantityUnit:           good.QuantityUnit,
			QuantityUnitAmount:     good.QuantityUnitAmount,
			TestOnly:               good.TestOnly,
			Posters:                good.Posters,
			Labels:                 good.Labels,
			BenefitIntervalHours:   good.BenefitIntervalHours,
			GoodTotal:              good.GoodTotal,
			GoodSpotQuantity:       good.GoodSpotQuantity,
			StartAt:                good.StartAt,
			StartMode:              good.StartMode,
			CreatedAt:              good.CreatedAt,
			SaleStartAt:            good.SaleStartAt,
			SaleEndAt:              good.SaleEndAt,
			ServiceStartAt:         good.ServiceStartAt,
			TechnicalFeeRatio:      good.TechnicalFeeRatio,
			ElectricityFeeRatio:    good.ElectricityFeeRatio,
			Descriptions:           good.Descriptions,
			GoodBanner:             good.GoodBanner,
			DisplayNames:           good.DisplayNames,
			EnablePurchase:         good.EnablePurchase,
			EnableProductPage:      good.EnableProductPage,
			CancelMode:             good.CancelMode,
			DisplayColors:          good.DisplayColors,
			CancellableBeforeStart: good.CancellableBeforeStart,
			ProductPage:            good.ProductPage,
			EnableSetCommission:    good.EnableSetCommission,
			Likes:                  good.Likes,
			Dislikes:               good.Dislikes,
			Score:                  good.Score,
			ScoreCount:             good.ScoreCount,
			RecommendCount:         good.RecommendCount,
			CommentCount:           good.CommentCount,
			AppGoodStockID:         good.AppGoodStockID,
			AppGoodReserved:        good.AppGoodReserved,
			AppSpotQuantity:        good.AppSpotQuantity,
			AppGoodLocked:          good.AppGoodLocked,
			AppGoodWaitStart:       good.AppGoodWaitStart,
			AppGoodInService:       good.AppGoodInService,
			AppGoodSold:            good.AppGoodSold,
			LastRewardAt:           good.LastRewardAt,
			LastRewardAmount:       good.LastRewardAmount,
			TotalRewardAmount:      good.TotalRewardAmount,
			LastUnitRewardAmount:   good.LastUnitRewardAmount,
			AppGoodPosters:         good.AppGoodPosters,
			UnitType:               good.UnitType,
			QuantityCalculateType:  good.QuantityCalculateType,
			DurationType:           good.DurationType,
			DurationCalculateType:  good.DurationCalculateType,
			MinOrderAmount:         good.MinOrderAmount,
			MaxOrderAmount:         good.MaxOrderAmount,
			MaxUserAmount:          good.MaxUserAmount,
			MinOrderDuration:       good.MinOrderDuration,
			MaxOrderDuration:       good.MaxOrderDuration,
			SettlementType:         good.SettlementType,
			PackageWithRequireds:   good.PackageWithRequireds,
		}

		coin, ok := h.appCoins[good.CoinTypeID]
		if ok {
			info.CoinLogo = coin.Logo
			info.CoinName = coin.Name
			info.CoinUnit = coin.Unit
			info.CoinPreSale = coin.Presale
			info.CoinEnv = coin.ENV
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	good, err := appgoodmwcli.GetGood(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if good == nil {
		return nil, fmt.Errorf("invalid appgood")
	}

	handler := &queryHandler{
		Handler:  h,
		goods:    []*appgoodmwpb.Good{good},
		appCoins: map[string]*appcoinmwpb.Coin{},
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetGoods(ctx context.Context) ([]*npool.Good, uint32, error) {
	conds := &appgoodmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	goods, total, err := appgoodmwcli.GetGoods(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(goods) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:  h,
		goods:    goods,
		appCoins: map[string]*appcoinmwpb.Coin{},
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, total, nil
	}

	return handler.infos, total, nil
}

func (h *Handler) GetGoodOnly(ctx context.Context) (*npool.Good, error) {
	conds := &appgoodmwpb.Conds{}
	if h.EntID != nil {
		conds.EntID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID}
	}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	good, err := appgoodmwcli.GetGoodOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	if good == nil {
		return nil, nil
	}

	handler := &queryHandler{
		Handler:  h,
		goods:    []*appgoodmwpb.Good{good},
		appCoins: map[string]*appcoinmwpb.Coin{},
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
