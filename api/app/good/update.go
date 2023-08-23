//nolint:dupl
package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
)

func (s *Server) UpdateAppGood(ctx context.Context, in *npool.UpdateAppGoodRequest) (*npool.UpdateAppGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithAppID(&in.AppID, true),
		good1.WithOnline(in.Online, false),
		good1.WithVisible(in.Visible, false),
		good1.WithGoodName(in.GoodName, false),
		good1.WithPrice(in.Price, false),
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
		good1.WithPosters(in.Posters, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateNAppGood(ctx context.Context, in *npool.UpdateNAppGoodRequest) (*npool.UpdateNAppGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithAppID(&in.TargetAppID, true),
		good1.WithOnline(in.Online, false),
		good1.WithVisible(in.Visible, false),
		good1.WithGoodName(in.GoodName, false),
		good1.WithPrice(in.Price, false),
		good1.WithDisplayIndex(in.DisplayIndex, false),
		good1.WithPurchaseLimit(in.PurchaseLimit, false),
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
		good1.WithPosters(in.Posters, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNAppGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNAppGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateNAppGoodResponse{
		Info: info,
	}, nil
}
