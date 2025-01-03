package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
)

func (s *Server) AdminUpdatePledge(ctx context.Context, in *npool.AdminUpdatePledgeRequest) (*npool.AdminUpdatePledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(&in.ID, true),
		pledge1.WithEntID(&in.EntID, true),
		pledge1.WithGoodID(&in.GoodID, true),

		pledge1.WithContractCodeURL(in.ContractCodeURL, false),
		pledge1.WithContractCodeBranch(in.ContractCodeBranch, false),
		pledge1.WithName(in.Name, false),
		pledge1.WithServiceStartAt(in.ServiceStartAt, false),
		pledge1.WithStartMode(in.StartMode, false),
		pledge1.WithTestOnly(in.TestOnly, false),
		pledge1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		pledge1.WithPurchasable(in.Purchasable, false),
		pledge1.WithOnline(in.Online, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdatePledgeResponse{
		Info: info,
	}, nil
}
