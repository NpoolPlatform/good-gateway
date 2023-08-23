//nolint:dupl
package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-gateway/pkg/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
)

func (s *Server) UpdateComment(ctx context.Context, in *npool.UpdateCommentRequest) (*npool.UpdateCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithAppID(&in.AppID, true),
		comment1.WithUserID(&in.UserID, true),
		comment1.WithContent(&in.Content, true),
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
