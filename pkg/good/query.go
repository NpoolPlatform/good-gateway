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
		coinTypeIDs = append(coinTypeIDs, good.SupportCoinTypeIDs...)
	}
	coins, _, err := coinmwcli.GetCoins(ctx, &coinmwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}, int32(0), int32(len(coinTypeIDs)))
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.coins[coin.ID] = coin
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, good := range h.goods {
		info := &npool.Good{
			ID:                     good.ID,
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
			GoodType:               good.GoodType,
			BenefitType:            good.BenefitType,
			Price:                  good.Price,
			Title:                  good.Title,
			Unit:                   good.Unit,
			UnitAmount:             good.UnitAmount,
			TestOnly:               good.TestOnly,
			Posters:                good.Posters,
			Labels:                 good.Labels,
			StockID:                good.GoodStockID,
			Total:                  good.GoodTotal,
			Locked:                 good.GoodLocked,
			InService:              good.GoodInService,
			WaitStart:              good.GoodWaitStart,
			Sold:                   good.GoodSold,
			DeliveryAt:             good.DeliveryAt,
			StartAt:                good.StartAt,
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
		}

		coin, ok := h.coins[good.CoinTypeID]
		if ok {
			info.CoinLogo = coin.Logo
			info.CoinName = coin.Name
			info.CoinUnit = coin.Unit
			info.CoinPreSale = coin.Presale
		}

		supportCoins := []*npool.Good_CoinInfo{}
		for _, coinTypeID := range good.SupportCoinTypeIDs {
			coin, ok := h.coins[coinTypeID]
			if !ok {
				continue
			}
			supportCoins = append(supportCoins, &npool.Good_CoinInfo{
				CoinTypeID:  coinTypeID,
				CoinLogo:    coin.Logo,
				CoinName:    coin.Name,
				CoinUnit:    coin.Unit,
				CoinPreSale: coin.Presale,
			})
		}

		info.SupportCoins = supportCoins
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	good, err := goodmwcli.GetGood(ctx, *h.ID)
	if err != nil {
		return nil, err
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
