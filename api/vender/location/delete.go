//nolint:dupl
package location

import (
	"context"

	location1 "github.com/NpoolPlatform/good-gateway/pkg/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/vender/location"
)

func (s *Server) DeleteLocation(ctx context.Context, in *npool.DeleteLocationRequest) (*npool.DeleteLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLocationResponse{
		Info: info,
	}, nil
}
