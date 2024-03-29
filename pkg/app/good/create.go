package good

import (
	"context"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
)

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := appgoodmwcli.CreateGood(ctx, &appgoodmwpb.GoodReq{
		EntID:                  h.EntID,
		AppID:                  h.AppID,
		GoodID:                 h.GoodID,
		Online:                 h.Online,
		Visible:                h.Visible,
		GoodName:               h.GoodName,
		UnitPrice:              h.UnitPrice,
		PackagePrice:           h.PackagePrice,
		DisplayIndex:           h.DisplayIndex,
		SaleStartAt:            h.SaleStartAt,
		SaleEndAt:              h.SaleEndAt,
		ServiceStartAt:         h.ServiceStartAt,
		TechnicalFeeRatio:      h.TechniqueFeeRatio,
		ElectricityFeeRatio:    h.ElectricityFeeRatio,
		Descriptions:           h.Descriptions,
		GoodBanner:             h.GoodBanner,
		DisplayNames:           h.DisplayNames,
		EnablePurchase:         h.EnablePurchase,
		EnableProductPage:      h.EnableProductPage,
		CancelMode:             h.CancelMode,
		DisplayColors:          h.DisplayColors,
		CancellableBeforeStart: h.CancellableBeforeStart,
		ProductPage:            h.ProductPage,
		EnableSetCommission:    h.EnableSetCommission,
		Posters:                h.Posters,
		MinOrderAmount:         h.MinOrderAmount,
		MaxOrderAmount:         h.MaxOrderAmount,
		MaxUserAmount:          h.MaxUserAmount,
		MinOrderDuration:       h.MinOrderDuration,
		MaxOrderDuration:       h.MaxOrderDuration,
		PackageWithRequireds:   h.PackageWithRequireds,
	}); err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
