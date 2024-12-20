package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-gateway/pkg/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
)

func (s *Server) AdminCreatePledge(ctx context.Context, in *npool.AdminCreatePledgeRequest) (*npool.AdminCreatePledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithContractCodeURL(&in.ContractCodeURL, true),
		pledge1.WithContractCodeBranch(&in.ContractCodeBranch, true),
		pledge1.WithCoinTypeID(&in.CoinTypeID, true),
		pledge1.WithGoodType(&in.GoodType, true),
		pledge1.WithName(&in.Name, true),
		pledge1.WithServiceStartAt(in.ServiceStartAt, true),
		pledge1.WithStartMode(&in.StartMode, true),
		pledge1.WithTestOnly(in.TestOnly, false),
		pledge1.WithBenefitIntervalHours(in.BenefitIntervalHours, true),
		pledge1.WithPurchasable(in.Purchasable, false),
		pledge1.WithOnline(in.Online, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreatePledgeResponse{
		Info: info,
	}, nil
}
