package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-gateway/pkg/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/fee"
)

func (s *Server) AdminUpdateFee(ctx context.Context, in *npool.AdminUpdateFeeRequest) (*npool.AdminUpdateFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithID(&in.ID, true),
		fee1.WithEntID(&in.EntID, true),
		fee1.WithGoodID(&in.GoodID, true),
		fee1.WithGoodType(in.GoodType, false),
		fee1.WithName(in.Name, false),
		fee1.WithSettlementType(in.SettlementType, false),
		fee1.WithUnitValue(in.UnitValue, false),
		fee1.WithDurationType(in.DurationType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateFeeResponse{
		Info: info,
	}, nil
}
