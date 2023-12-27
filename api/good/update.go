package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
)

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(&in.ID, true),
		good1.WithEntID(&in.EntID, true),
		good1.WithDeviceInfoID(in.DeviceInfoID, false),
		good1.WithDurationDays(in.DurationDays, false),
		good1.WithCoinTypeID(in.CoinTypeID, false),
		good1.WithVendorLocationID(in.VendorLocationID, false),
		good1.WithUnitPrice(in.UnitPrice, false),
		good1.WithBenefitType(in.BenefitType, false),
		good1.WithGoodType(in.GoodType, false),
		good1.WithTitle(in.Title, false),
		good1.WithQuantityUnit(in.QuantityUnit, false),
		good1.WithQuantityUnitAmount(in.QuantityUnitAmount, false),
		good1.WithDeliveryAt(in.DeliveryAt, false),
		good1.WithStartAt(in.StartAt, false),
		good1.WithStartMode(in.StartMode, false),
		good1.WithTestOnly(in.TestOnly, false),
		good1.WithTotal(in.Total, false),
		good1.WithPosters(in.Posters, false),
		good1.WithLabels(in.Labels, false),
		good1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		good1.WithUnitLockDeposit(in.UnitLockDeposit, false),
		good1.WithUnitType(in.UnitType, false),
		good1.WithQuantityCalculateType(in.QuantityCalculateType, false),
		good1.WithDurationType(in.DurationType, false),
		good1.WithDurationCalculateType(in.DurationCalculateType, false),
		good1.WithSettlementType(in.SettlementType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}
