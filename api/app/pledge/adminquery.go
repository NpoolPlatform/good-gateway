package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

func (s *Server) AdminGetAppPledges(ctx context.Context, in *npool.AdminGetAppPledgesRequest) (*npool.AdminGetAppPledgesResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithAppID(&in.TargetAppID, true),
		pledge1.WithOffset(in.Offset),
		pledge1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppPledges",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPledges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppPledges",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppPledgesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
