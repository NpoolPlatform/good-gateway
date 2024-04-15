package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
)

func (s *Server) GetRecommends(ctx context.Context, in *npool.GetRecommendsRequest) (*npool.GetRecommendsResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithAppGoodID(in.AppGoodID, false),
		recommend1.WithRecommenderID(in.UserID, false),
		recommend1.WithOffset(in.Offset),
		recommend1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRecommends(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRecommendsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
