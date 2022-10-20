package good

import (
	"context"

	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"

	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"
)

func GetGood(ctx context.Context, id string) (*npool.Good, error) {
	info, err := goodmwcli.GetGood(ctx, id)
	if err != nil {
		return nil, err
	}
	return ScanCoinType(ctx, info)
}

func GetGoods(ctx context.Context, offset, limit int32) ([]*npool.Good, uint32, error) {
	infos, total, err := goodmwcli.GetGoods(ctx, nil, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	goods := []*npool.Good{}
	for _, val := range infos {
		good, err := ScanCoinType(ctx, val)
		if err != nil {
			return nil, 0, err
		}

		goods = append(goods, good)
	}
	return goods, total, nil
}

func CreateGood(ctx context.Context, req *npool.CreateGoodRequest) (*npool.Good, error) {
	locked := int32(req.Locked)
	inService := int32(req.InService)
	info, err := goodmwcli.CreateGood(ctx, &goodmwpb.GoodReq{
		DeviceInfoID:       &req.DeviceInfoID,
		DurationDays:       &req.DurationDays,
		CoinTypeID:         &req.CoinTypeID,
		InheritFromGoodID:  &req.InheritFromGoodID,
		VendorLocationID:   &req.VendorLocationID,
		Price:              &req.Price,
		BenefitType:        &req.BenefitType,
		GoodType:           &req.GoodType,
		Title:              &req.Title,
		Unit:               &req.Unit,
		UnitAmount:         &req.UnitAmount,
		SupportCoinTypeIDs: req.SupportCoinTypeIDs,
		DeliveryAt:         &req.DeliveryAt,
		StartAt:            &req.StartAt,
		TestOnly:           &req.TestOnly,
		Total:              &req.Total,
		Locked:             &locked,
		InService:          &inService,
		Sold:               &req.Sold,
		Posters:            req.Posters,
		Labels:             req.Labels,
	})
	if err != nil {
		return nil, err
	}
	return ScanCoinType(ctx, info)
}

func UpdateGood(ctx context.Context, req *npool.UpdateGoodRequest) (*npool.Good, error) {
	info, err := goodmwcli.UpdateGood(ctx, &goodmwpb.GoodReq{
		ID:                 &req.ID,
		DeviceInfoID:       req.DeviceInfoID,
		DurationDays:       req.DurationDays,
		CoinTypeID:         req.CoinTypeID,
		InheritFromGoodID:  req.InheritFromGoodID,
		VendorLocationID:   req.VendorLocationID,
		Price:              req.Price,
		Title:              req.Title,
		Unit:               req.Unit,
		UnitAmount:         req.UnitAmount,
		SupportCoinTypeIDs: req.SupportCoinTypeIDs,
		DeliveryAt:         req.DeliveryAt,
		StartAt:            req.StartAt,
		TestOnly:           req.TestOnly,
		Total:              req.Total,
		Sold:               req.Sold,
		Posters:            req.Posters,
		Labels:             req.Labels,
	})
	if err != nil {
		return nil, err
	}
	return ScanCoinType(ctx, info)
}

func ScanCoinType(ctx context.Context, info *goodmwpb.Good) (*npool.Good, error) {
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

	return &npool.Good{
		ID:                         info.ID,
		DeviceInfoID:               info.DeviceInfoID,
		DeviceType:                 info.DeviceType,
		DeviceManufacturer:         info.DeviceManufacturer,
		DevicePowerComsuption:      info.DevicePowerComsuption,
		DeviceShipmentAt:           info.DeviceShipmentAt,
		DevicePosters:              info.DevicePosters,
		DurationDays:               info.DurationDays,
		CoinTypeID:                 info.CoinTypeID,
		CoinLogo:                   coinType.Logo,
		CoinName:                   coinType.Name,
		CoinUnit:                   coinType.Unit,
		CoinPreSale:                coinType.PreSale,
		InheritFromGoodID:          info.InheritFromGoodID,
		InheritFromGoodName:        info.InheritFromGoodName,
		InheritFromGoodType:        info.InheritFromGoodType,
		InheritFromGoodBenefitType: info.InheritFromGoodBenefitType,
		VendorLocationID:           info.VendorLocationID,
		VendorLocationCountry:      info.VendorLocationCountry,
		VendorLocationProvince:     info.VendorLocationProvince,
		VendorLocationCity:         info.VendorLocationCity,
		VendorLocationAddress:      info.VendorLocationAddress,
		GoodType:                   info.GoodType,
		BenefitType:                info.BenefitType,
		Price:                      info.Price,
		Title:                      info.Title,
		Unit:                       info.Unit,
		UnitAmount:                 info.UnitAmount,
		TestOnly:                   info.TestOnly,
		Posters:                    info.Posters,
		Labels:                     info.Labels,
		VoteCount:                  info.VoteCount,
		Rating:                     info.Rating,
		SupportCoins:               supportCoins,
		GoodStockID:                info.GoodStockID,
		GoodTotal:                  info.GoodTotal,
		GoodLocked:                 info.GoodLocked,
		GoodInService:              info.GoodInService,
		GoodSold:                   info.GoodSold,
		DeliveryAt:                 info.DeliveryAt,
		StartAt:                    info.StartAt,
		CreatedAt:                  info.CreatedAt,
		UpdatedAt:                  info.UpdatedAt,
	}, nil
}
