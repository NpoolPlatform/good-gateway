package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-gateway/pkg/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device/manufacturer"
)

func (s *Server) AdminUpdateManufacturer(ctx context.Context, in *npool.AdminUpdateManufacturerRequest) (*npool.AdminUpdateManufacturerResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithID(&in.ID, true),
		manufacturer1.WithEntID(&in.EntID, true),
		manufacturer1.WithName(in.Name, false),
		manufacturer1.WithLogo(in.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateManufacturer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateManufacturerResponse{
		Info: info,
	}, nil
}
