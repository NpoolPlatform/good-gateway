package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
)

func (s *Server) UpdateSimulate(ctx context.Context, in *npool.UpdateSimulateRequest) (*npool.UpdateSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithID(&in.ID, true),
		simulate1.WithEntID(&in.EntID, true),
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithOrderUnits(in.OrderUnits, false),
		simulate1.WithOrderDuration(in.OrderDuration, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSimulateResponse{
		Info: info,
	}, nil
}
