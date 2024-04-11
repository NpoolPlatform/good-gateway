package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/good-gateway/pkg/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device"
)

func (s *Server) CreateDeviceType(ctx context.Context, in *npool.CreateDeviceTypeRequest) (*npool.CreateDeviceTypeResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithType(&in.Type, true),
		devicetype1.WithManufacturer(&in.Manufacturer, true),
		devicetype1.WithPowerConsumption(&in.PowerConsumption, true),
		devicetype1.WithShipmentAt(&in.ShipmentAt, true),
		devicetype1.WithPosters(in.Posters, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDeviceTypeResponse{
		Info: info,
	}, nil
}
