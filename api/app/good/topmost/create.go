package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

func (s *Server) CreateTopMost(ctx context.Context, in *npool.CreateTopMostRequest) (*npool.CreateTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithAppID(&in.AppID, true),
		topmost1.WithTopMostType(&in.TopMostType, true),
		topmost1.WithTitle(&in.Title, true),
		topmost1.WithMessage(&in.Message, true),
		topmost1.WithTargetURL(in.TargetUrl, true),
		topmost1.WithStartAt(&in.StartAt, true),
		topmost1.WithEndAt(&in.EndAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostResponse{
		Info: info,
	}, nil
}
