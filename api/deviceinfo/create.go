package deviceinfo

import (
	"context"

	deviceinfo1 "github.com/NpoolPlatform/good-gateway/pkg/deviceinfo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/deviceinfo"
)

func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithType(&in.Type, true),
		deviceinfo1.WithManufacturer(&in.Manufacturer, true),
		deviceinfo1.WithPowerConsumption(&in.PowerConsumption, true),
		deviceinfo1.WithShipmentAt(&in.ShipmentAt, true),
		deviceinfo1.WithPosters(in.Posters, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDeviceInfoResponse{
		Info: info,
	}, nil
}
