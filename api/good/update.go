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
		good1.WithDeviceInfoID(in.DeviceInfoID, false),
		good1.WithDurationDays(in.DurationDays, false),
		good1.WithCoinTypeID(in.CoinTypeID, false),
		good1.WithVendorLocationID(in.VendorLocationID, false),
		good1.WithPrice(in.Price, false),
		good1.WithBenefitType(in.BenefitType, false),
		good1.WithGoodType(in.GoodType, false),
		good1.WithTitle(in.Title, false),
		good1.WithUnit(in.Unit, false),
		good1.WithUnitAmount(in.UnitAmount, false),
		good1.WithSupportCoinTypeIDs(in.SupportCoinTypeIDs, false),
		good1.WithDeliveryAt(in.DeliveryAt, false),
		good1.WithStartAt(in.StartAt, false),
		good1.WithTestOnly(in.TestOnly, false),
		good1.WithTotal(in.Total, false),
		good1.WithPosters(in.Posters, false),
		good1.WithLabels(in.Labels, false),
		good1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		good1.WithUnitLockDeposit(in.UnitLockDeposit, false),
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
