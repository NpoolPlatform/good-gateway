package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
)

func (s *Server) GetAppPledge(ctx context.Context, in *npool.GetAppPledgeRequest) (*npool.GetAppPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppPledgeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppPledges(ctx context.Context, in *npool.GetAppPledgesRequest) (*npool.GetAppPledgesResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithAppID(&in.AppID, true),
		pledge1.WithOffset(in.Offset),
		pledge1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPledges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppPledgesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
