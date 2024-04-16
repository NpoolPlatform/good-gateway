//nolint:dupl
package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/required"
)

func (s *Server) AdminCreateRequired(ctx context.Context, in *npool.AdminCreateRequiredRequest) (*npool.AdminCreateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithAppID(&in.TargetAppID, true),
		required1.WithMainAppGoodID(&in.MainAppGoodID, true),
		required1.WithRequiredAppGoodID(&in.RequiredAppGoodID, true),
		required1.WithMust(in.Must, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateRequiredResponse{
		Info: info,
	}, nil
}
