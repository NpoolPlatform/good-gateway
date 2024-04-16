package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/required"
)

func (s *Server) GetRequireds(ctx context.Context, in *npool.GetRequiredsRequest) (*npool.GetRequiredsResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithAppID(&in.AppID, true),
		required1.WithAppGoodID(in.AppGoodID, false),
		required1.WithOffset(in.Offset),
		required1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRequireds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRequiredsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
