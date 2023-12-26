package good

import (
	"context"
	"fmt"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
)

func (h *Handler) UpdateGood(ctx context.Context) (*npool.Good, error) {
	exist, err := appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid appgood")
	}

	if _, err := appgoodmwcli.UpdateGood(ctx, &appgoodmwpb.GoodReq{
		ID:                     h.ID,
		Online:                 h.Online,
		Visible:                h.Visible,
		GoodName:               h.GoodName,
		UnitPrice:              h.UnitPrice,
		PackagePrice:           h.PackagePrice,
		DisplayIndex:           h.DisplayIndex,
		PurchaseLimit:          h.PurchaseLimit,
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
		UserPurchaseLimit:      h.UserPurchaseLimit,
		DisplayColors:          h.DisplayColors,
		CancellableBeforeStart: h.CancellableBeforeStart,
		ProductPage:            h.ProductPage,
		EnableSetCommission:    h.EnableSetCommission,
		Posters:                h.Posters,
	}); err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
