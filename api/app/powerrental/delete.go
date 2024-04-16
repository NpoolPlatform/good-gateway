package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-gateway/pkg/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental"
)

func (s *Server) DeleteAppPowerRental(ctx context.Context, in *npool.DeleteAppPowerRentalRequest) (*npool.DeleteAppPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAppPowerRentalResponse{
		Info: info,
	}, nil
}
