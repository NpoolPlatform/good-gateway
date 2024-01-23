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

func (s *Server) CreateTopMostGood(ctx context.Context, in *npool.CreateTopMostGoodRequest) (*npool.CreateTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithAppID(&in.AppID, true),
		topmostgood1.WithTopMostID(&in.TopMostID, true),
		topmostgood1.WithAppGoodID(&in.AppGoodID, true),
		topmostgood1.WithPosters(in.Posters, true),
		topmostgood1.WithUnitPrice(in.UnitPrice, false),
		topmostgood1.WithPackagePrice(in.PackagePrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateNTopMostGood(ctx context.Context, in *npool.CreateNTopMostGoodRequest) (*npool.CreateNTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithAppID(&in.TargetAppID, true),
		topmostgood1.WithTopMostID(&in.TopMostID, true),
		topmostgood1.WithAppGoodID(&in.AppGoodID, true),
		topmostgood1.WithPosters(in.Posters, true),
		topmostgood1.WithUnitPrice(in.UnitPrice, false),
		topmostgood1.WithPackagePrice(in.PackagePrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateNTopMostGoodResponse{
		Info: info,
	}, nil
}
