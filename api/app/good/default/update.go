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

func (s *Server) UpdateDefault(ctx context.Context, in *npool.UpdateDefaultRequest) (*npool.UpdateDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithID(&in.ID, true),
		default1.WithEntID(&in.EntID, true),
		default1.WithAppID(&in.AppID, true),
		default1.WithAppGoodID(in.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDefaultResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateNDefault(ctx context.Context, in *npool.UpdateNDefaultRequest) (*npool.UpdateNDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithID(&in.ID, true),
		default1.WithEntID(&in.EntID, true),
		default1.WithAppID(&in.TargetAppID, true),
		default1.WithAppGoodID(in.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateNDefaultResponse{
		Info: info,
	}, nil
}
