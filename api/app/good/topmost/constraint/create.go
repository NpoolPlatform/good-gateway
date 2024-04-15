//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/constraint"
)

func (s *Server) CreateTopMostConstraint(ctx context.Context, in *npool.CreateTopMostConstraintRequest) (*npool.CreateTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithTopMostID(&in.TopMostID, true),
		constraint1.WithConstraint(&in.Constraint, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostConstraintResponse{
		Info: info,
	}, nil
}
