//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"
)

func (s *Server) AdminUpdateTopMostGoodConstraint(ctx context.Context, in *npool.AdminUpdateTopMostGoodConstraintRequest) (*npool.AdminUpdateTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.TargetAppID, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
