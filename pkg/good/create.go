//nolint:dupl
package good

import (
	"context"

	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/google/uuid"
)

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := goodmwcli.CreateGood(ctx, &goodmwpb.GoodReq{
		EntID:                 h.EntID,
		DeviceInfoID:          h.DeviceInfoID,
		DurationDays:          h.DurationDays,
		CoinTypeID:            h.CoinTypeID,
		VendorLocationID:      h.VendorLocationID,
		UnitPrice:             h.UnitPrice,
		BenefitType:           h.BenefitType,
		GoodType:              h.GoodType,
		Title:                 h.Title,
		QuantityUnit:          h.QuantityUnit,
		QuantityUnitAmount:    h.QuantityUnitAmount,
		DeliveryAt:            h.DeliveryAt,
		StartAt:               h.StartAt,
		StartMode:             h.StartMode,
		TestOnly:              h.TestOnly,
		Total:                 h.Total,
		Posters:               h.Posters,
		Labels:                h.Labels,
		BenefitIntervalHours:  h.BenefitIntervalHours,
		UnitLockDeposit:       h.UnitLockDeposit,
		UnitType:              h.UnitType,
		QuantityCalculateType: h.QuantityCalculateType,
		DurationType:          h.DurationType,
		DurationCalculateType: h.DurationCalculateType,
	}); err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
