package good

import (
	"context"
	"fmt"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

type queryHandler struct {
	*Handler
	goods []*goodmwpb.Good
	infos []*npool.Good
	coins map[string]*coinmwpb.Coin
}

func (h *queryHandler) getCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, good := range h.goods {
		coinTypeIDs = append(coinTypeIDs, good.CoinTypeID)
	}
	coins, _, err := coinmwcli.GetCoins(ctx, &coinmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}, int32(0), int32(len(coinTypeIDs)))
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.coins[coin.EntID] = coin
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, good := range h.goods {
		info := &npool.Good{
			ID:                     good.ID,
			EntID:                  good.EntID,
			DeviceInfoID:           good.DeviceInfoID,
			DeviceType:             good.DeviceType,
			DeviceManufacturer:     good.DeviceManufacturer,
			DevicePowerConsumption: good.DevicePowerConsumption,
			DeviceShipmentAt:       good.DeviceShipmentAt,
			DevicePosters:          good.DevicePosters,
			DurationDays:           good.DurationDays,
			CoinTypeID:             good.CoinTypeID,
			VendorLocationID:       good.VendorLocationID,
			VendorLocationCountry:  good.VendorLocationCountry,
			VendorLocationProvince: good.VendorLocationProvince,
			VendorLocationCity:     good.VendorLocationCity,
			VendorLocationAddress:  good.VendorLocationAddress,
			VendorBrandName:        good.VendorBrandName,
			VendorBrandLogo:        good.VendorBrandLogo,
			GoodType:               good.GoodType,
			BenefitType:            good.BenefitType,
			Price:                  good.Price,
			Title:                  good.Title,
			QuantityUnit:           good.QuantityUnit,
			QuantityUnitAmount:     good.QuantityUnitAmount,
			TestOnly:               good.TestOnly,
			Posters:                good.Posters,
			Labels:                 good.Labels,
			StockID:                good.GoodStockID,
			Total:                  good.GoodTotal,
			SpotQuantity:           good.GoodSpotQuantity,
			Locked:                 good.GoodLocked,
			InService:              good.GoodInService,
			WaitStart:              good.GoodWaitStart,
			Sold:                   good.GoodSold,
			AppReserved:            good.GoodAppReserved,
			DeliveryAt:             good.DeliveryAt,
			StartAt:                good.StartAt,
			StartMode:              good.StartMode,
			CreatedAt:              good.CreatedAt,
			UpdatedAt:              good.UpdatedAt,
			BenefitIntervalHours:   good.BenefitIntervalHours,
			Likes:                  good.Likes,
			Dislikes:               good.Dislikes,
			Score:                  good.Score,
			ScoreCount:             good.ScoreCount,
			RecommendCount:         good.RecommendCount,
			CommentCount:           good.CommentCount,
			UnitLockDeposit:        good.UnitLockDeposit,
			UnitType:               good.UnitType,
			QuantityCalculateType:  good.QuantityCalculateType,
			DurationType:           good.DurationType,
			DurationCalculateType:  good.DurationCalculateType,
		}

		coin, ok := h.coins[good.CoinTypeID]
		if ok {
			info.CoinLogo = coin.Logo
			info.CoinName = coin.Name
			info.CoinUnit = coin.Unit
			info.CoinPreSale = coin.Presale
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	good, err := goodmwcli.GetGood(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if good == nil {
		return nil, fmt.Errorf("invalid good")
	}

	handler := &queryHandler{
		Handler: h,
		goods:   []*goodmwpb.Good{good},
		coins:   map[string]*coinmwpb.Coin{},
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
	goods, total, err := goodmwcli.GetGoods(ctx, &goodmwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(goods) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		goods:   goods,
		coins:   map[string]*coinmwpb.Coin{},
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
