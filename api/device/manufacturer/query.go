package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-gateway/pkg/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/device/manufacturer"
)

func (s *Server) GetManufacturers(ctx context.Context, in *npool.GetManufacturersRequest) (*npool.GetManufacturersResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithOffset(in.Offset),
		manufacturer1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturers",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetManufacturers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturers",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetManufacturersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
