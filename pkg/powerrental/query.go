package powerrental

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	goodcoingwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
)

type queryHandler struct {
	*Handler
	powerRentals []*powerrentalmwpb.PowerRental
	coins        map[string]*coinmwpb.Coin
	infos        []*npool.PowerRental
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = goodgwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, powerRental := range h.powerRentals {
			for _, goodCoin := range powerRental.GoodCoins {
				coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, powerRental := range h.powerRentals {
		info := &npool.PowerRental{
			ID:     powerRental.ID,
			EntID:  powerRental.EntID,
			GoodID: powerRental.GoodID,

			DeviceTypeID:           powerRental.DeviceTypeID,
			DeviceType:             powerRental.DeviceType,
			DeviceManufacturerName: powerRental.DeviceManufacturerName,
			DeviceManufacturerLogo: powerRental.DeviceManufacturerLogo,
			DevicePowerConsumption: powerRental.DevicePowerConsumption,
			DeviceShipmentAt:       powerRental.DeviceShipmentAt,

			VendorLocationID: powerRental.VendorLocationID,
			VendorBrand:      powerRental.VendorBrand,
			VendorLogo:       powerRental.VendorLogo,
			VendorCountry:    powerRental.VendorCountry,
			VendorProvince:   powerRental.VendorProvince,

			UnitPrice:           powerRental.UnitPrice,
			QuantityUnit:        powerRental.QuantityUnit,
			QuantityUnitAmount:  powerRental.QuantityUnitAmount,
			DeliveryAt:          powerRental.DeliveryAt,
			UnitLockDeposit:     powerRental.UnitLockDeposit,
			DurationDisplayType: powerRental.DurationDisplayType,

			GoodType:             powerRental.GoodType,
			BenefitType:          powerRental.BenefitType,
			Name:                 powerRental.Name,
			ServiceStartAt:       powerRental.ServiceStartAt,
			StartMode:            powerRental.StartMode,
			TestOnly:             powerRental.TestOnly,
			BenefitIntervalHours: powerRental.BenefitIntervalHours,
			Purchasable:          powerRental.Purchasable,
			Online:               powerRental.Online,

			StockMode:    powerRental.StockMode,
			StockID:      powerRental.GoodStockID,
			Total:        powerRental.GoodTotal,
			SpotQuantity: powerRental.GoodSpotQuantity,
			Locked:       powerRental.GoodLocked,
			InService:    powerRental.GoodInService,
			WaitStart:    powerRental.GoodWaitStart,
			Sold:         powerRental.GoodSold,
			AppReserved:  powerRental.GoodAppReserved,

			RewardState:  powerRental.RewardState,
			LastRewardAt: powerRental.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range powerRental.Rewards {
					coin, ok := h.coins[reward.CoinTypeID]
					if !ok {
						continue
					}
					rewards = append(rewards, &goodcoinrewardgwpb.RewardInfo{
						CoinTypeID:            reward.CoinTypeID,
						CoinName:              coin.Name,
						CoinUnit:              coin.Unit,
						CoinENV:               coin.ENV,
						CoinLogo:              coin.Logo,
						RewardTID:             reward.RewardTID,
						NextRewardStartAmount: reward.NextRewardStartAmount,
						LastRewardAmount:      reward.LastRewardAmount,
						LastUnitRewardAmount:  reward.LastUnitRewardAmount,
						TotalRewardAmount:     reward.TotalRewardAmount,
						MainCoin:              reward.MainCoin,
					})
				}
				return
			}(),

			GoodCoins: func() (coins []*goodcoingwpb.GoodCoinInfo) {
				for _, goodCoin := range powerRental.GoodCoins {
					coin, ok := h.coins[goodCoin.CoinTypeID]
					if !ok {
						continue
					}
					coins = append(coins, &goodcoingwpb.GoodCoinInfo{
						CoinTypeID: goodCoin.CoinTypeID,
						CoinName:   coin.Name,
						CoinUnit:   coin.Unit,
						CoinENV:    coin.ENV,
						CoinLogo:   coin.Logo,
						Main:       goodCoin.Main,
						Index:      goodCoin.Index,
					})
				}
				return
			}(),
			MiningGoodStocks: powerRental.MiningGoodStocks,

			CreatedAt: powerRental.CreatedAt,
			UpdatedAt: powerRental.UpdatedAt,
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.PowerRental, error) {
	powerRental, err := powerrentalmwcli.GetPowerRental(ctx, *h.GoodID)
	if err != nil {
		return nil, err
	}
	if powerRental == nil {
		return nil, fmt.Errorf("invalid powerrental")
	}

	handler := &queryHandler{
		Handler:      h,
		powerRentals: []*powerrentalmwpb.PowerRental{powerRental},
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

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.PowerRental, uint32, error) {
	powerRentals, total, err := powerrentalmwcli.GetPowerRentals(ctx, &powerrentalmwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(powerRentals) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:      h,
		powerRentals: powerRentals,
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
