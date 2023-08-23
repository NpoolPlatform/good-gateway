//nolint:dupl
package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
)

func (s *Server) GetGood(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithAppID(&in.AppID, true),
		good1.WithOffset(in.Offset),
		good1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetNGoods(ctx context.Context, in *npool.GetNGoodsRequest) (*npool.GetNGoodsResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithAppID(&in.TargetAppID, true),
		good1.WithOffset(in.Offset),
		good1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetNGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetNGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
