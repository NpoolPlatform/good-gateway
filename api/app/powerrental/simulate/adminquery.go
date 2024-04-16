package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
)

func (s *Server) AdminGetSimulates(ctx context.Context, in *npool.AdminGetSimulatesRequest) (*npool.AdminGetSimulatesResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.TargetAppID, true),
		simulate1.WithOffset(in.Offset),
		simulate1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSimulates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetSimulatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
