package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"
)

func (s *Server) AdminDeleteTopMostGoodConstraint(ctx context.Context, in *npool.AdminDeleteTopMostGoodConstraintRequest) (*npool.AdminDeleteTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
