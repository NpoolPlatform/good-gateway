//nolint:dupl
package goodcoin

import (
	"context"

	goodcoin1 "github.com/NpoolPlatform/good-gateway/pkg/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
)

func (s *Server) AdminCreateGoodCoin(ctx context.Context, in *npool.AdminCreateGoodCoinRequest) (*npool.AdminCreateGoodCoinResponse, error) {
	handler, err := goodcoin1.NewHandler(
		ctx,
		goodcoin1.WithGoodID(&in.GoodID, true),
		goodcoin1.WithCoinTypeID(&in.CoinTypeID, true),
		goodcoin1.WithMain(in.Main, false),
		goodcoin1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateGoodCoinResponse{
		Info: info,
	}, nil
}
