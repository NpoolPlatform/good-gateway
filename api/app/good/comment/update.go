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
func (s *Server) UpdateComment(ctx context.Context, in *npool.UpdateCommentRequest) (*npool.UpdateCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.AppID, true),
		comment1.WithCommentUserID(&in.UserID, true),
		comment1.WithContent(in.Content, false),
		comment1.WithAnonymous(in.Anonymous, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCommentResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) UpdateUserComment(ctx context.Context, in *npool.UpdateUserCommentRequest) (*npool.UpdateUserCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.AppID, true),
		comment1.WithCommentUserID(&in.TargetUserID, true),
		comment1.WithHide(in.Hide, false),
		comment1.WithHideReason(in.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateUserCommentResponse{
		Info: info,
	}, nil
}
