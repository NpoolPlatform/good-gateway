package powerrental

import (
	"context"

	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"

	"github.com/google/uuid"
)

func (h *Handler) CreatePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
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
		MinOrderDuration:             h.MinOrderDuration,
		MaxOrderDuration:             h.MaxOrderDuration,
		UnitPrice:                    h.UnitPrice,
		SaleStartAt:                  h.SaleStartAt,
		SaleEndAt:                    h.SaleEndAt,
		SaleMode:                     h.SaleMode,
		FixedDuration:                h.FixedDuration,
		PackageWithRequireds:         h.PackageWithRequireds,
	}); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}