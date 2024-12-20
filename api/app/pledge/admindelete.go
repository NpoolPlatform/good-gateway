package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

func (s *Server) AdminDeleteAppPledge(ctx context.Context, in *npool.AdminDeleteAppPledgeRequest) (*npool.AdminDeleteAppPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(&in.ID, true),
		pledge1.WithEntID(&in.EntID, true),
		pledge1.WithAppID(&in.TargetAppID, true),
		pledge1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppPledgeResponse{
		Info: info,
	}, nil
}
