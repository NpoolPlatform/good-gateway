package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

func (s *Server) AdminUpdateAppPledge(ctx context.Context, in *npool.AdminUpdateAppPledgeRequest) (*npool.AdminUpdateAppPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(&in.ID, true),
		pledge1.WithEntID(&in.EntID, true),
		pledge1.WithAppID(&in.TargetAppID, true),
		pledge1.WithAppGoodID(&in.AppGoodID, true),

		pledge1.WithPurchasable(in.Purchasable, false),
		pledge1.WithEnableProductPage(in.EnableProductPage, false),
		pledge1.WithProductPage(in.ProductPage, false),
		pledge1.WithOnline(in.Online, false),
		pledge1.WithVisible(in.Visible, false),
		pledge1.WithName(in.Name, false),
		pledge1.WithDisplayIndex(in.DisplayIndex, false),
		pledge1.WithBanner(in.Banner, false),

		pledge1.WithServiceStartAt(in.ServiceStartAt, false),
		pledge1.WithStartMode(in.StartMode, false),
		pledge1.WithEnableSetCommission(in.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateAppPledgeResponse{
		Info: info,
	}, nil
}
