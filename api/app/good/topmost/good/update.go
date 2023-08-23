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

func (s *Server) UpdateTopMostGood(ctx context.Context, in *npool.UpdateTopMostGoodRequest) (*npool.UpdateTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithAppID(&in.AppID, true),
		topmostgood1.WithPosters(in.Posters, true),
		topmostgood1.WithPrice(in.Price, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateNTopMostGood(ctx context.Context, in *npool.UpdateNTopMostGoodRequest) (*npool.UpdateNTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithAppID(&in.TargetAppID, true),
		topmostgood1.WithPosters(in.Posters, true),
		topmostgood1.WithPrice(in.Price, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateNTopMostGoodResponse{
		Info: info,
	}, nil
}
