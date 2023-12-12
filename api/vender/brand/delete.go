//nolint:dupl
package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/good-gateway/pkg/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/vender/brand"
)

func (s *Server) DeleteBrand(ctx context.Context, in *npool.DeleteBrandRequest) (*npool.DeleteBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(&in.ID, true),
		brand1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteBrandResponse{
		Info: info,
	}, nil
}
