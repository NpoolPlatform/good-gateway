package powerrental

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
)

type queryHandler struct {
	*Handler
	appPowerRentals []*apppowerrentalmwpb.PowerRental
	infos           []*npool.AppPowerRental
	apps            map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appPowerRental := range h.appPowerRentals {
			appIDs = append(appIDs, appPowerRental.AppID)
		}
		return
	}())
	return err
}

//nolint:funlen
func (h *queryHandler) formalize() {
	for _, appPowerRental := range h.appPowerRentals {
		info := &npool.AppPowerRental{
			ID:        appPowerRental.ID,
			EntID:     appPowerRental.EntID,
			AppID:     appPowerRental.AppID,
			GoodID:    appPowerRental.GoodID,
			AppGoodID: appPowerRental.AppGoodID,

			DeviceTypeID:           appPowerRental.DeviceTypeID,
			DeviceType:             appPowerRental.DeviceType,
			DeviceManufacturerName: appPowerRental.DeviceManufacturerName,
			DeviceManufacturerLogo: appPowerRental.DeviceManufacturerLogo,
			DevicePowerConsumption: appPowerRental.DevicePowerConsumption,
			DeviceShipmentAt:       appPowerRental.DeviceShipmentAt,

			VendorLocationID: appPowerRental.VendorLocationID,
			VendorBrand:      appPowerRental.VendorBrand,
			VendorLogo:       appPowerRental.VendorLogo,
			VendorCountry:    appPowerRental.VendorCountry,
			VendorProvince:   appPowerRental.VendorProvince,

			UnitPrice:          appPowerRental.UnitPrice,
			QuantityUnit:       appPowerRental.QuantityUnit,
			QuantityUnitAmount: appPowerRental.QuantityUnitAmount,
			DeliveryAt:         appPowerRental.DeliveryAt,
			UnitLockDeposit:    appPowerRental.UnitLockDeposit,
			DurationType:       appPowerRental.DurationType,

			GoodType:             appPowerRental.GoodType,
			BenefitType:          appPowerRental.BenefitType,
			GoodName:             appPowerRental.GoodName,
			ServiceStartAt:       appPowerRental.ServiceStartAt,
			StartMode:            appPowerRental.StartMode,
			TestOnly:             appPowerRental.TestOnly,
			BenefitIntervalHours: appPowerRental.BenefitIntervalHours,
			GoodPurchasable:      appPowerRental.GoodPurchasable,
			GoodOnline:           appPowerRental.GoodOnline,

			StockMode:                    appPowerRental.StockMode,
			AppGoodPurchasable:           appPowerRental.AppGoodPurchasable,
			AppGoodOnline:                appPowerRental.AppGoodOnline,
			EnableProductPage:            appPowerRental.EnableProductPage,
			ProductPage:                  appPowerRental.ProductPage,
			Visible:                      appPowerRental.Visible,
			AppGoodName:                  appPowerRental.AppGoodName,
			DisplayIndex:                 appPowerRental.DisplayIndex,
			Banner:                       appPowerRental.Banner,
			CancelMode:                   appPowerRental.CancelMode,
			CancelableBeforeStartSeconds: appPowerRental.CancelableBeforeStartSeconds,
			EnableSetCommission:          appPowerRental.EnableSetCommission,
			MinOrderAmount:               appPowerRental.MinOrderAmount,
			MaxOrderAmount:               appPowerRental.MaxOrderAmount,
			MaxUserAmount:                appPowerRental.MaxUserAmount,
			MinOrderDuration:             appPowerRental.MinOrderDuration,
			MaxOrderDuration:             appPowerRental.MaxOrderDuration,
			SaleStartAt:                  appPowerRental.SaleStartAt,
			SaleEndAt:                    appPowerRental.SaleEndAt,
			SaleMode:                     appPowerRental.SaleMode,
			FixedDuration:                appPowerRental.FixedDuration,
			PackageWithRequireds:         appPowerRental.PackageWithRequireds,

			GoodStockID:      appPowerRental.GoodStockID,
			GoodTotal:        appPowerRental.GoodTotal,
			GoodSpotQuantity: appPowerRental.GoodSpotQuantity,

			AppGoodStockID:      appPowerRental.AppGoodStockID,
			AppGoodReserved:     appPowerRental.AppGoodReserved,
			AppGoodSpotQuantity: appPowerRental.AppGoodSpotQuantity,
			AppGoodLocked:       appPowerRental.AppGoodLocked,
			AppGoodInService:    appPowerRental.AppGoodInService,
			AppGoodWaitStart:    appPowerRental.AppGoodWaitStart,
			AppGoodSold:         appPowerRental.AppGoodSold,

			Likes:          appPowerRental.Likes,
			Dislikes:       appPowerRental.Dislikes,
			Score:          appPowerRental.Score,
			ScoreCount:     appPowerRental.ScoreCount,
			RecommendCount: appPowerRental.RecommendCount,
			CommentCount:   appPowerRental.CommentCount,

			LastRewardAt:         appPowerRental.LastRewardAt,
			LastRewardAmount:     appPowerRental.LastRewardAmount,
			TotalRewardAmount:    appPowerRental.TotalRewardAmount,
			LastUnitRewardAmount: appPowerRental.LastUnitRewardAmount,

			// TODO: expand coin information
			GoodCoins:     appPowerRental.GoodCoins,
			Descriptions:  appPowerRental.Descriptions,
			Posters:       appPowerRental.Posters,
			DisplayNames:  appPowerRental.DisplayNames,
			DisplayColors: appPowerRental.DisplayColors,
			// TODO: expand mining pool information
			AppMiningGoodStocks: appPowerRental.AppMiningGoodStocks,
			MiningGoodStocks:    appPowerRental.MiningGoodStocks,
			Labels:              appPowerRental.Labels,

			CreatedAt: appPowerRental.CreatedAt,
			UpdatedAt: appPowerRental.UpdatedAt,
		}
		if appPowerRental.GoodType == types.GoodType_LegacyPowerRental {
			info.TechniqueFeeRatio = &appPowerRental.TechniqueFeeRatio
		}
		app, ok := h.apps[appPowerRental.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	appPowerRental, err := apppowerrentalmwcli.GetPowerRental(ctx, *h.AppGoodID)
	if err != nil {
		return nil, err
	}
	if appPowerRental == nil {
		return nil, fmt.Errorf("invalid apppowerrental")
	}

	handler := &queryHandler{
		Handler:         h,
		appPowerRentals: []*apppowerrentalmwpb.PowerRental{appPowerRental},
		apps:            map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.AppPowerRental, uint32, error) {
	appPowerRentals, total, err := apppowerrentalmwcli.GetPowerRentals(ctx, &apppowerrentalmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(appPowerRentals) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:         h,
		appPowerRentals: appPowerRentals,
		apps:            map[string]*appmwpb.App{},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
