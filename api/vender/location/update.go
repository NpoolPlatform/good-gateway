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

func (s *Server) UpdateLocation(ctx context.Context, in *npool.UpdateLocationRequest) (*npool.UpdateLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithID(&in.ID, true),
		location1.WithCountry(in.Country, false),
		location1.WithProvince(in.Province, false),
		location1.WithCity(in.City, false),
		location1.WithAddress(in.Address, false),
		location1.WithBrandID(in.BrandID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateLocationResponse{
		Info: info,
	}, nil
}
