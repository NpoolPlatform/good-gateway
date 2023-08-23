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

func (s *Server) GetDefaults(ctx context.Context, in *npool.GetDefaultsRequest) (*npool.GetDefaultsResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithAppID(&in.AppID, true),
		default1.WithOffset(in.Offset),
		default1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDefaults(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDefaultsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetNDefaults(ctx context.Context, in *npool.GetNDefaultsRequest) (*npool.GetNDefaultsResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithAppID(&in.TargetAppID, true),
		default1.WithOffset(in.Offset),
		default1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetNDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDefaults(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetNDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetNDefaultsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
