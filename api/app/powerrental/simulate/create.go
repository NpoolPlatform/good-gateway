package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
)

func (s *Server) CreateSimulate(ctx context.Context, in *npool.CreateSimulateRequest) (*npool.CreateSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithAppGoodID(&in.AppGoodID, true),
		simulate1.WithOrderUnits(&in.OrderUnits, true),
		simulate1.WithOrderDurationSeconds(&in.OrderDurationSeconds, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSimulateResponse{
		Info: info,
	}, nil
}
