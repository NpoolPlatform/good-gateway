package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/good-gateway/pkg/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/delegatedstaking"
)

func (s *Server) AdminDeleteDelegatedStaking(ctx context.Context, in *npool.AdminDeleteDelegatedStakingRequest) (*npool.AdminDeleteDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(&in.ID, true),
		delegatedstaking1.WithEntID(&in.EntID, true),
		delegatedstaking1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteDelegatedStakingResponse{
		Info: info,
	}, nil
}
