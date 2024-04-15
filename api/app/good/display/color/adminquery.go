package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/color"
)

func (s *Server) AdminGetDisplayColors(ctx context.Context, in *npool.AdminGetDisplayColorsRequest) (*npool.AdminGetDisplayColorsResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithAppID(&in.TargetAppID, true),
		displaycolor1.WithOffset(in.Offset),
		displaycolor1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayColors(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetDisplayColorsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
