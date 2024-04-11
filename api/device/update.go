package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/good-gateway/pkg/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device"
)

func (s *Server) UpdateDeviceType(ctx context.Context, in *npool.UpdateDeviceTypeRequest) (*npool.UpdateDeviceTypeResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithID(&in.ID, true),
		devicetype1.WithEntID(&in.EntID, true),
		devicetype1.WithType(in.Type, false),
		devicetype1.WithManufacturer(in.Manufacturer, false),
		devicetype1.WithPowerConsumption(in.PowerConsumption, false),
		devicetype1.WithShipmentAt(in.ShipmentAt, false),
		devicetype1.WithPosters(in.Posters, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDeviceTypeResponse{
		Info: info,
	}, nil
}
