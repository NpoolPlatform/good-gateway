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

func (s *Server) AdminGetComments(ctx context.Context, in *npool.AdminGetCommentsRequest) (*npool.AdminGetCommentsResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithAppID(&in.TargetAppID, true),
		comment1.WithCommentUserID(in.TargetUserID, false),
		comment1.WithAppGoodID(in.AppGoodID, false),
		comment1.WithOffset(in.Offset),
		comment1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetComments",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetComments(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetComments",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetCommentsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
