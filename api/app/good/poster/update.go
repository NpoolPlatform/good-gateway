//nolint:dupl
package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/poster"
)

func (s *Server) UpdatePoster(ctx context.Context, in *npool.UpdatePosterRequest) (*npool.UpdatePosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(&in.ID, true),
		poster1.WithEntID(&in.EntID, true),
		poster1.WithAppID(&in.AppID, true),
		poster1.WithPoster(in.Poster, false),
		poster1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdatePosterResponse{
		Info: info,
	}, nil
}
