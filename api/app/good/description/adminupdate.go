//nolint:dupl
package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
)

func (s *Server) AdminUpdateDescription(ctx context.Context, in *npool.AdminUpdateDescriptionRequest) (*npool.AdminUpdateDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID, true),
		description1.WithEntID(&in.EntID, true),
		description1.WithAppID(&in.TargetAppID, true),
		description1.WithDescription(in.Description, false),
		description1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDescriptionResponse{
		Info: info,
	}, nil
}
