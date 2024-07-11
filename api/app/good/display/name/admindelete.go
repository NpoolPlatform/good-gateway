package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/name"
)

func (s *Server) AdminDeleteDisplayName(ctx context.Context, in *npool.AdminDeleteDisplayNameRequest) (*npool.AdminDeleteDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithID(&in.ID, true),
		displayname1.WithEntID(&in.EntID, true),
		displayname1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteDisplayNameResponse{
		Info: info,
	}, nil
}
