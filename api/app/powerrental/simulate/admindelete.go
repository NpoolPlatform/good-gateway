package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
)

func (s *Server) AdminDeleteSimulate(ctx context.Context, in *npool.AdminDeleteSimulateRequest) (*npool.AdminDeleteSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithID(&in.ID, true),
		simulate1.WithEntID(&in.EntID, true),
		simulate1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteSimulateResponse{
		Info: info,
	}, nil
}
