//nolint:dupl
package appgood

import (
	"context"

	appgood1 "github.com/NpoolPlatform/good-gateway/pkg/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
)

func (s *Server) AdminGetGoods(ctx context.Context, in *npool.AdminGetGoodsRequest) (*npool.AdminGetGoodsResponse, error) {
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithAppID(&in.TargetAppID, true),
		appgood1.WithOffset(in.Offset),
		appgood1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
