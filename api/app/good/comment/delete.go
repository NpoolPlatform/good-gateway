//nolint:dupl
package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/comment"
)

func (s *Server) DeleteComment(ctx context.Context, in *npool.DeleteCommentRequest) (*npool.DeleteCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.AppID, true),
		comment1.WithCommentUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCommentResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteUserComment(ctx context.Context, in *npool.DeleteUserCommentRequest) (*npool.DeleteUserCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.AppID, true),
		comment1.WithCommentUserID(&in.TargetUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserCommentResponse{
		Info: info,
	}, nil
}
