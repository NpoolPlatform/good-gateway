package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
)

func (s *Server) AdminUpdateDefault(ctx context.Context, in *npool.AdminUpdateDefaultRequest) (*npool.AdminUpdateDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithID(&in.ID, true),
		default1.WithEntID(&in.EntID, true),
		default1.WithAppID(&in.TargetAppID, true),
		default1.WithAppGoodID(in.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDefaultResponse{
		Info: info,
	}, nil
}
