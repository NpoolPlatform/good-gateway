package goodcoin

import (
	"context"

	goodcoin1 "github.com/NpoolPlatform/good-gateway/pkg/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
)

func (s *Server) GetGoodCoins(ctx context.Context, in *npool.GetGoodCoinsRequest) (*npool.GetGoodCoinsResponse, error) {
	handler, err := goodcoin1.NewHandler(
		ctx,
		goodcoin1.WithOffset(in.Offset),
		goodcoin1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoodCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
