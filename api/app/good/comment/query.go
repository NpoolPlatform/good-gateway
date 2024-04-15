package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/comment"
)

func (s *Server) GetComments(ctx context.Context, in *npool.GetCommentsRequest) (*npool.GetCommentsResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithAppID(&in.AppID, true),
		comment1.WithAppGoodID(in.AppGoodID, false),
		comment1.WithUserID(in.UserID, false),
		comment1.WithOffset(in.Offset),
		comment1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComments",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetComments(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComments",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCommentsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
