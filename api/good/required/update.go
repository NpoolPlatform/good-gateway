package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-gateway/pkg/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/required"
)

func (s *Server) UpdateRequired(ctx context.Context, in *npool.UpdateRequiredRequest) (*npool.UpdateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(&in.ID, true),
		required1.WithMust(&in.Must, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRequiredResponse{
		Info: info,
	}, nil
}
