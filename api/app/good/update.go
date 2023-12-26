package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
)

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithEntID(&in.EntID, true),
		good1.WithAppID(&in.AppID, true),
		good1.WithOnline(in.Online, false),
		good1.WithVisible(in.Visible, false),
		good1.WithGoodName(in.GoodName, false),
		good1.WithUnitPrice(in.UnitPrice, false),
		good1.WithPackagePrice(in.PackagePrice, false),
		good1.WithDisplayIndex(in.DisplayIndex, false),
		good1.WithPurchaseLimit(in.PurchaseLimit, false),
		good1.WithSaleStartAt(in.SaleStartAt, false),
		good1.WithSaleEndAt(in.SaleEndAt, false),
		good1.WithServiceStartAt(in.ServiceStartAt, false),
		good1.WithTechniqueFeeRatio(in.TechnicalFeeRatio, false),
		good1.WithElectricityFeeRatio(in.ElectricityFeeRatio, false),
		good1.WithEnablePurchase(in.EnablePurchase, false),
		good1.WithEnableProductPage(in.EnableProductPage, false),
		good1.WithCancelMode(in.CancelMode, false),
		good1.WithUserPurchaseLimit(in.UserPurchaseLimit, false),
		good1.WithDisplayColors(in.DisplayColors, false),
		good1.WithCancellableBeforeStart(in.CancellableBeforeStart, false),
		good1.WithProductPage(in.ProductPage, false),
		good1.WithEnableSetCommission(in.EnableSetCommission, false),
		good1.WithPosters(in.Posters, false),
		good1.WithDescriptions(in.Descriptions, false),
		good1.WithGoodBanner(in.GoodBanner, false),
		good1.WithDisplayNames(in.DisplayNames, false),
		good1.WithMinOrderAmount(in.MinOrderAmount, false),
		good1.WithMaxOrderAmount(in.MaxOrderAmount, false),
		good1.WithMaxUserAmount(in.MaxUserAmount, false),
		good1.WithMinOrderDuration(in.MinOrderDuration, false),
		good1.WithMaxOrderDuration(in.MaxOrderDuration, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateNGood(ctx context.Context, in *npool.UpdateNGoodRequest) (*npool.UpdateNGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithEntID(&in.EntID, true),
		good1.WithAppID(&in.TargetAppID, true),
		good1.WithOnline(in.Online, false),
		good1.WithVisible(in.Visible, false),
		good1.WithGoodName(in.GoodName, false),
		good1.WithUnitPrice(in.UnitPrice, false),
		good1.WithPackagePrice(in.PackagePrice, false),
		good1.WithDisplayIndex(in.DisplayIndex, false),
		good1.WithPurchaseLimit(in.PurchaseLimit, false),
		good1.WithTechniqueFeeRatio(in.TechnicalFeeRatio, false),
		good1.WithElectricityFeeRatio(in.ElectricityFeeRatio, false),
		good1.WithEnablePurchase(in.EnablePurchase, false),
		good1.WithEnableProductPage(in.EnableProductPage, false),
		good1.WithCancelMode(in.CancelMode, false),
		good1.WithUserPurchaseLimit(in.UserPurchaseLimit, false),
		good1.WithDisplayColors(in.DisplayColors, false),
		good1.WithCancellableBeforeStart(in.CancellableBeforeStart, false),
		good1.WithProductPage(in.ProductPage, false),
		good1.WithEnableSetCommission(in.EnableSetCommission, false),
		good1.WithPosters(in.Posters, false),
		good1.WithDescriptions(in.Descriptions, false),
		good1.WithGoodBanner(in.GoodBanner, false),
		good1.WithDisplayNames(in.DisplayNames, false),
		good1.WithServiceStartAt(in.ServiceStartAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateNGoodResponse{
		Info: info,
	}, nil
}
