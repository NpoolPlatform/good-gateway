package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
)

func (s *Server) CreateDescription(ctx context.Context, in *npool.CreateDescriptionRequest) (*npool.CreateDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.AppID, true),
		description1.WithAppGoodID(&in.AppGoodID, true),
		description1.WithDescription(&in.Description, true),
		description1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDescriptionResponse{
		Info: info,
	}, nil
}
