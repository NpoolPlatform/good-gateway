//nolint:dupl
package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
)

func (s *Server) AdminUpdateTopMostGood(ctx context.Context, in *npool.AdminUpdateTopMostGoodRequest) (*npool.AdminUpdateTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithEntID(&in.EntID, true),
		topmostgood1.WithAppID(&in.TargetAppID, true),
		topmostgood1.WithUnitPrice(in.UnitPrice, false),
		topmostgood1.WithDisplayIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateTopMostGoodResponse{
		Info: info,
	}, nil
}
