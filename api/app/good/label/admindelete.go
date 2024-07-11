package label

import (
	"context"

	label1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/label"
)

func (s *Server) AdminDeleteLabel(ctx context.Context, in *npool.AdminDeleteLabelRequest) (*npool.AdminDeleteLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithID(&in.ID, true),
		label1.WithEntID(&in.EntID, true),
		label1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteLabelResponse{
		Info: info,
	}, nil
}
