package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-gateway/pkg/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"
)

func (s *Server) GetPowerRental(ctx context.Context, in *npool.GetPowerRentalRequest) (*npool.GetPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPowerRentalResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPowerRentals(ctx context.Context, in *npool.GetPowerRentalsRequest) (*npool.GetPowerRentalsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPowerRentals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPowerRentalsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
