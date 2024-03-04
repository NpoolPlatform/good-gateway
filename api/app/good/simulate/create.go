//nolint:dupl
package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/simulate"
)

func (s *Server) CreateSimulate(ctx context.Context, in *npool.CreateSimulateRequest) (*npool.CreateSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithAppGoodID(&in.AppGoodID, true),
		simulate1.WithFixedOrderUnits(&in.FixedOrderUnits, true),
		simulate1.WithFixedOrderDuration(&in.FixedOrderDuration, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSimulateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateNSimulate(ctx context.Context, in *npool.CreateNSimulateRequest) (*npool.CreateNSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.TargetAppID, true),
		simulate1.WithAppGoodID(&in.AppGoodID, true),
		simulate1.WithFixedOrderUnits(&in.FixedOrderUnits, true),
		simulate1.WithFixedOrderDuration(&in.FixedOrderDuration, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateNSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateNSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateNSimulateResponse{
		Info: info,
	}, nil
}
