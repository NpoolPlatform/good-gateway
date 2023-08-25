package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
)

func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithAppID(&in.TargetAppID, true),
		good1.WithGoodID(&in.GoodID, true),
		good1.WithOnline(&in.Online, true),
		good1.WithVisible(&in.Visible, true),
		good1.WithGoodName(&in.GoodName, true),
		good1.WithPrice(&in.Price, true),
		good1.WithDisplayIndex(&in.DisplayIndex, true),
		good1.WithPurchaseLimit(&in.PurchaseLimit, true),
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
		good1.WithDescriptions(in.Descriptions, false),
		good1.WithGoodBanner(in.GoodBanner, false),
		good1.WithDisplayNames(in.DisplayNames, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodResponse{
		Info: info,
	}, nil
}
