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

func (s *Server) GetDeviceInfos(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithOffset(in.Offset),
		deviceinfo1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfos",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDeviceInfos(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfos",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceInfosResponse{
		Infos: infos,
		Total: total,
	}, nil
}
