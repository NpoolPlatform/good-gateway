//nolint:dupl
package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/simulate"
)

func (s *Server) GetSimulates(ctx context.Context, in *npool.GetSimulatesRequest) (*npool.GetSimulatesResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithOffset(in.Offset),
		simulate1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSimulates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSimulatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetNSimulates(ctx context.Context, in *npool.GetNSimulatesRequest) (*npool.GetNSimulatesResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.TargetAppID, true),
		simulate1.WithOffset(in.Offset),
		simulate1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetNSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSimulates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetNSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNSimulatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
