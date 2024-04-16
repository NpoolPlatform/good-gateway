package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-gateway/pkg/device/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device/poster"
)

func (s *Server) AdminUpdatePoster(ctx context.Context, in *npool.AdminUpdatePosterRequest) (*npool.AdminUpdatePosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(&in.ID, true),
		poster1.WithEntID(&in.EntID, true),
		poster1.WithPoster(in.Poster, false),
		poster1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdatePosterResponse{
		Info: info,
	}, nil
}
