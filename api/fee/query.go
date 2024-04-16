package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-gateway/pkg/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/fee"
)

func (s *Server) GetFees(ctx context.Context, in *npool.GetFeesRequest) (*npool.GetFeesResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithOffset(in.Offset),
		fee1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetFees(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFeesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
