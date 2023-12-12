//nolint:dupl
package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-gateway/pkg/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/required"
)

func (s *Server) CreateRequired(ctx context.Context, in *npool.CreateRequiredRequest) (*npool.CreateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithMainGoodID(&in.MainGoodID, true),
		required1.WithRequiredGoodID(&in.RequiredGoodID, true),
		required1.WithMust(&in.Must, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRequiredResponse{
		Info: info,
	}, nil
}
