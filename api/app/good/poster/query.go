package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/poster"
)

func (s *Server) GetPosters(ctx context.Context, in *npool.GetPostersRequest) (*npool.GetPostersResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithAppID(&in.AppID, true),
		poster1.WithAppGoodID(in.AppGoodID, false),
		poster1.WithOffset(in.Offset),
		poster1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPosters(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPostersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
