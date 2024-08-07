package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

func (s *Server) GetTopMosts(ctx context.Context, in *npool.GetTopMostsRequest) (*npool.GetTopMostsResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithAppID(&in.AppID, true),
		topmost1.WithOffset(in.Offset),
		topmost1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMosts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
