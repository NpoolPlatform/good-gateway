package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

func (s *Server) AdminCreateAppPledge(ctx context.Context, in *npool.AdminCreateAppPledgeRequest) (*npool.AdminCreateAppPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithAppID(&in.TargetAppID, true),
		pledge1.WithGoodID(&in.GoodID, true),

		pledge1.WithPurchasable(in.Purchasable, false),
		pledge1.WithEnableProductPage(in.EnableProductPage, false),
		pledge1.WithProductPage(in.ProductPage, false),
		pledge1.WithOnline(in.Online, false),
		pledge1.WithVisible(in.Visible, false),
		pledge1.WithName(&in.Name, true),
		pledge1.WithDisplayIndex(in.DisplayIndex, false),
		pledge1.WithBanner(in.Banner, false),

		pledge1.WithServiceStartAt(&in.ServiceStartAt, true),
		pledge1.WithStartMode(in.StartMode, false),
		pledge1.WithEnableSetCommission(in.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppPledgeResponse{
		Info: info,
	}, nil
}
