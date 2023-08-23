package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

func (s *Server) DeleteTopMost(ctx context.Context, in *npool.DeleteTopMostRequest) (*npool.DeleteTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
		topmost1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostResponse{
		Info: info,
	}, nil
}
