package deviceinfo

import (
	"context"

	deviceinfo1 "github.com/NpoolPlatform/good-gateway/pkg/deviceinfo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/deviceinfo"
)

func (s *Server) DeleteDeviceInfo(ctx context.Context, in *npool.DeleteDeviceInfoRequest) (*npool.DeleteDeviceInfoResponse, error) {
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithID(&in.ID, true),
		deviceinfo1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDeviceInfoResponse{
		Info: info,
	}, nil
}
