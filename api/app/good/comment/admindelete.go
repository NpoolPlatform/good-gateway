package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/comment"
)

//nolint
func (s *Server) AdminDeleteComment(ctx context.Context, in *npool.AdminDeleteCommentRequest) (*npool.AdminDeleteCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.TargetAppID, true),
		comment1.WithTargetUserID(&in.TargetUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteCommentResponse{
		Info: info,
	}, nil
}
