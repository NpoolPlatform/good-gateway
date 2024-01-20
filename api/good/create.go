package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-gateway/pkg/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
)

func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithDeviceInfoID(&in.DeviceInfoID, true),
		good1.WithCoinTypeID(&in.CoinTypeID, true),
		good1.WithVendorLocationID(&in.VendorLocationID, true),
		good1.WithUnitPrice(&in.UnitPrice, true),
		good1.WithBenefitType(&in.BenefitType, true),
		good1.WithGoodType(&in.GoodType, true),
		good1.WithTitle(&in.Title, true),
		good1.WithQuantityUnit(&in.QuantityUnit, true),
		good1.WithQuantityUnitAmount(&in.QuantityUnitAmount, true),
		good1.WithDeliveryAt(&in.DeliveryAt, true),
		good1.WithStartAt(&in.StartAt, true),
		good1.WithStartMode(&in.StartMode, true),
		good1.WithTestOnly(&in.TestOnly, true),
		good1.WithTotal(&in.Total, true),
		good1.WithPosters(in.Posters, true),
		good1.WithLabels(in.Labels, true),
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
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodResponse{
		Info: info,
	}, nil
}
