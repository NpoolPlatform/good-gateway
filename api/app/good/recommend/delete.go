package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
)

func (s *Server) DeleteRecommend(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(&in.ID, true),
		recommend1.WithEntID(&in.EntID, true),
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithRecommenderID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRecommendResponse{
		Info: info,
	}, nil
}
