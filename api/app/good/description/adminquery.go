package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
)

func (s *Server) AdminGetDescriptions(ctx context.Context, in *npool.AdminGetDescriptionsRequest) (*npool.AdminGetDescriptionsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.TargetAppID, true),
		description1.WithOffset(in.Offset),
		description1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDescriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetDescriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
