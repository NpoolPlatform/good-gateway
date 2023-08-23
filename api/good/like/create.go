//nolint:dupl
package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-gateway/pkg/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/like"
)

func (s *Server) CreateLike(ctx context.Context, in *npool.CreateLikeRequest) (*npool.CreateLikeResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(&in.UserID, true),
		like1.WithGoodID(&in.GoodID, true),
		like1.WithLike(&in.Like, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLike(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLikeResponse{
		Info: info,
	}, nil
}
