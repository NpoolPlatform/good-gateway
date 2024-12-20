package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
)

func (s *Server) GetPledge(ctx context.Context, in *npool.GetPledgeRequest) (*npool.GetPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPledgeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPledges(ctx context.Context, in *npool.GetPledgesRequest) (*npool.GetPledgesResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithOffset(in.Offset),
		pledge1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPledges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPledgesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
