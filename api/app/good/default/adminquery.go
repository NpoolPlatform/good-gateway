package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
)

func (s *Server) AdminGetDefaults(ctx context.Context, in *npool.AdminGetDefaultsRequest) (*npool.AdminGetDefaultsResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithAppID(&in.TargetAppID, true),
		default1.WithOffset(in.Offset),
		default1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDefaults(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetDefaultsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
