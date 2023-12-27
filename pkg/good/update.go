//nolint:dupl
package good

import (
	"context"
	"fmt"

	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (h *Handler) UpdateGood(ctx context.Context) (*npool.Good, error) {
	info, err := goodmwcli.GetGoodOnly(ctx, &goodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid good")
	}

	// TODO: if start mode is set from TBD to confirmed, we should update all order's start at and start mode

	if _, err := goodmwcli.UpdateGood(ctx, &goodmwpb.GoodReq{
		ID:                    h.ID,
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
		SettlementType:        h.SettlementType,
	}); err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
