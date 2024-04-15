package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-gateway/pkg/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device/manufacturer"
)

func (s *Server) AdminDeleteManufacturer(ctx context.Context, in *npool.AdminDeleteManufacturerRequest) (*npool.AdminDeleteManufacturerResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithID(&in.ID, true),
		manufacturer1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteManufacturer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteManufacturerResponse{
		Info: info,
	}, nil
}
