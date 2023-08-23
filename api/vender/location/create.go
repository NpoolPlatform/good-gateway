//nolint:dupl
package location

import (
	"context"

	location1 "github.com/NpoolPlatform/good-gateway/pkg/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/vender/location"
)

func (s *Server) CreateLocation(ctx context.Context, in *npool.CreateLocationRequest) (*npool.CreateLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithCountry(&in.Country, true),
		location1.WithProvince(&in.Province, true),
		location1.WithCity(&in.City, true),
		location1.WithAddress(&in.Address, true),
		location1.WithBrandID(&in.BrandID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLocationResponse{
		Info: info,
	}, nil
}
