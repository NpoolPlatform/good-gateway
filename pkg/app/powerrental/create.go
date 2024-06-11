package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"

	"github.com/google/uuid"
)

// TODO: check start mode with power rental start mode

func (h *Handler) CreatePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.FixedDuration != nil && *h.FixedDuration {
		if h.MaxOrderDurationSeconds != nil && *h.MinOrderDurationSeconds != *h.MaxOrderDurationSeconds {
			return nil, wlog.Errorf("invalid maxorderdurationseconds")
		}
		h.MaxOrderDurationSeconds = h.MinOrderDurationSeconds
	}
	if err := apppowerrentalmwcli.CreatePowerRental(ctx, &apppowerrentalmwpb.PowerRentalReq{
		EntID:                        h.EntID,
		AppID:                        h.AppID,
		GoodID:                       h.GoodID,
		AppGoodID:                    h.AppGoodID,
		Purchasable:                  h.Purchasable,
		EnableProductPage:            h.EnableProductPage,
		ProductPage:                  h.ProductPage,
		Online:                       h.Online,
		Visible:                      h.Visible,
		Name:                         h.Name,
		DisplayIndex:                 h.DisplayIndex,
		Banner:                       h.Banner,
		ServiceStartAt:               h.ServiceStartAt,
		CancelMode:                   h.CancelMode,
		CancelableBeforeStartSeconds: h.CancelableBeforeStartSeconds,
		EnableSetCommission:          h.EnableSetCommission,
		MinOrderAmount:               h.MinOrderAmount,
		MaxOrderAmount:               h.MaxOrderAmount,
		MaxUserAmount:                h.MaxUserAmount,
		MinOrderDurationSeconds:      h.MinOrderDurationSeconds,
		MaxOrderDurationSeconds:      h.MaxOrderDurationSeconds,
		UnitPrice:                    h.UnitPrice,
		SaleStartAt:                  h.SaleStartAt,
		SaleEndAt:                    h.SaleEndAt,
		SaleMode:                     h.SaleMode,
		FixedDuration:                h.FixedDuration,
		PackageWithRequireds:         h.PackageWithRequireds,
		StartMode:                    h.StartMode,
	}); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
