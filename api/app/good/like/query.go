package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/like"
)

func (s *Server) GetLikes(ctx context.Context, in *npool.GetLikesRequest) (*npool.GetLikesResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(in.UserID, false),
		like1.WithAppGoodID(in.AppGoodID, false),
		like1.WithOffset(in.Offset),
		like1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLikes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLikesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
