package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/color"
)

func (s *Server) AdminCreateDisplayColor(ctx context.Context, in *npool.AdminCreateDisplayColorRequest) (*npool.AdminCreateDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithAppID(&in.TargetAppID, true),
		displaycolor1.WithAppGoodID(&in.AppGoodID, true),
		displaycolor1.WithColor(&in.Color, true),
		displaycolor1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDisplayColorResponse{
		Info: info,
	}, nil
}
