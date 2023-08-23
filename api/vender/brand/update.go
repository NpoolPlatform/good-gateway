package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/good-gateway/pkg/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/vender/brand"
)

func (s *Server) UpdateBrand(ctx context.Context, in *npool.UpdateBrandRequest) (*npool.UpdateBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(&in.ID, true),
		brand1.WithName(in.Name, false),
		brand1.WithLogo(in.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateBrandResponse{
		Info: info,
	}, nil
}
