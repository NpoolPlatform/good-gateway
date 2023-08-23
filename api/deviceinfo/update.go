//nolint:dupl
package deviceinfo

import (
	"context"

	deviceinfo1 "github.com/NpoolPlatform/good-gateway/pkg/deviceinfo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/deviceinfo"
)

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithID(&in.ID, true),
		deviceinfo1.WithType(in.Type, false),
		deviceinfo1.WithManufacturer(in.Manufacturer, false),
		deviceinfo1.WithPowerConsumption(in.PowerConsumption, false),
		deviceinfo1.WithShipmentAt(in.ShipmentAt, false),
		deviceinfo1.WithPosters(in.Posters, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDeviceInfoResponse{
		Info: info,
	}, nil
}
