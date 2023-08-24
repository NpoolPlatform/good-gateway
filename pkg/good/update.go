//nolint:dupl
package good

import (
	"context"
	"fmt"

	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (h *Handler) UpdateGood(ctx context.Context) (*npool.Good, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: if start mode is set from TBD to confirmed, we should update all order's start at and start mode

	if _, err := goodmwcli.UpdateGood(ctx, &goodmwpb.GoodReq{
		ID:                   h.ID,
		DeviceInfoID:         h.DeviceInfoID,
		DurationDays:         h.DurationDays,
		CoinTypeID:           h.CoinTypeID,
		VendorLocationID:     h.VendorLocationID,
		Price:                h.Price,
		BenefitType:          h.BenefitType,
		GoodType:             h.GoodType,
		Title:                h.Title,
		Unit:                 h.Unit,
		UnitAmount:           h.UnitAmount,
		SupportCoinTypeIDs:   h.SupportCoinTypeIDs,
		DeliveryAt:           h.DeliveryAt,
		StartAt:              h.StartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		Total:                h.Total,
		Posters:              h.Posters,
		Labels:               h.Labels,
		BenefitIntervalHours: h.BenefitIntervalHours,
		UnitLockDeposit:      h.UnitLockDeposit,
	}); err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
