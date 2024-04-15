//nolint:dupl
package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
)

func (s *Server) AdminCreateDefault(ctx context.Context, in *npool.AdminCreateDefaultRequest) (*npool.AdminCreateDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithAppID(&in.TargetAppID, true),
		default1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDefaultResponse{
		Info: info,
	}, nil
}
