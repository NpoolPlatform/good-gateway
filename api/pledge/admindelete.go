package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
)

func (s *Server) AdminDeletePledge(ctx context.Context, in *npool.AdminDeletePledgeRequest) (*npool.AdminDeletePledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(&in.ID, true),
		pledge1.WithEntID(&in.EntID, true),
		pledge1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeletePledgeResponse{
		Info: info,
	}, nil
}
