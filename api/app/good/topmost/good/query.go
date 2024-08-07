package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
)

func (s *Server) GetTopMostGoods(ctx context.Context, in *npool.GetTopMostGoodsRequest) (*npool.GetTopMostGoodsResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithAppID(&in.AppID, true),
		topmostgood1.WithOffset(in.Offset),
		topmostgood1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMostGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
