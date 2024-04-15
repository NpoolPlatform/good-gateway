package label

import (
	"context"

	label1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/label"
)

func (s *Server) GetLabels(ctx context.Context, in *npool.GetLabelsRequest) (*npool.GetLabelsResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithAppID(&in.AppID, true),
		label1.WithAppGoodID(in.AppGoodID, false),
		label1.WithOffset(in.Offset),
		label1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabels",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLabels(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabels",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLabelsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
