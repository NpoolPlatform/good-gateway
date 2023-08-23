//nolint:dupl
package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-gateway/pkg/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/recommend"
)

func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithRecommenderID(&in.UserID, true),
		recommend1.WithGoodID(&in.GoodID, true),
		recommend1.WithRecommendIndex(&in.RecommendIndex, true),
		recommend1.WithMessage(&in.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: info,
	}, nil
}
