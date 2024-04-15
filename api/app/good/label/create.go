package label

import (
	"context"

	label1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/label"
)

func (s *Server) CreateLabel(ctx context.Context, in *npool.CreateLabelRequest) (*npool.CreateLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithAppID(&in.AppID, true),
		label1.WithAppGoodID(&in.AppGoodID, true),
		label1.WithIcon(in.Icon, false),
		label1.WithIconBgColor(in.IconBgColor, false),
		label1.WithLabel(&in.Label, true),
		label1.WithLabelBgColor(in.LabelBgColor, false),
		label1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLabelResponse{
		Info: info,
	}, nil
}
