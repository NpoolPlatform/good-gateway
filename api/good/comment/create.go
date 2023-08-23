package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-gateway/pkg/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
)

func (s *Server) CreateComment(ctx context.Context, in *npool.CreateCommentRequest) (*npool.CreateCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithAppID(&in.AppID, true),
		comment1.WithUserID(&in.UserID, true),
		comment1.WithGoodID(&in.GoodID, true),
		comment1.WithOrderID(in.OrderID, false),
		comment1.WithContent(&in.Content, true),
		comment1.WithReplyToID(in.ReplyToID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCommentResponse{
		Info: info,
	}, nil
}
