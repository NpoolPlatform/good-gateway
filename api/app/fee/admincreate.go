package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/good-gateway/pkg/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
)

func (s *Server) AdminCreateAppFee(ctx context.Context, in *npool.AdminCreateAppFeeRequest) (*npool.AdminCreateAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithAppID(&in.TargetAppID, true),
		appfee1.WithGoodID(&in.GoodID, true),
		appfee1.WithProductPage(in.ProductPage, false),
		appfee1.WithName(&in.Name, true),
		appfee1.WithBanner(in.Banner, false),
		appfee1.WithUnitValue(&in.UnitValue, true),
		appfee1.WithMinOrderDuration(&in.MinOrderDuration, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppFeeResponse{
		Info: info,
	}, nil
}
