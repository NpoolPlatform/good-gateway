package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/good-gateway/pkg/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device"
)

func (s *Server) AdminDeleteDeviceType(ctx context.Context, in *npool.AdminDeleteDeviceTypeRequest) (*npool.AdminDeleteDeviceTypeResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithID(&in.ID, true),
		devicetype1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteDeviceTypeResponse{
		Info: info,
	}, nil
}
