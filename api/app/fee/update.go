package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/good-gateway/pkg/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
)

func (s *Server) UpdateAppFee(ctx context.Context, in *npool.UpdateAppFeeRequest) (*npool.UpdateAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithID(&in.ID, true),
		appfee1.WithEntID(&in.EntID, true),
		appfee1.WithAppID(&in.AppID, true),
		appfee1.WithAppGoodID(&in.AppGoodID, true),
		appfee1.WithProductPage(in.ProductPage, false),
		appfee1.WithName(in.Name, false),
		appfee1.WithBanner(in.Banner, false),
		appfee1.WithUnitValue(in.UnitValue, false),
		appfee1.WithMinOrderDurationSeconds(in.MinOrderDurationSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppFeeResponse{
		Info: info,
	}, nil
}
