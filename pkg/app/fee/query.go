package appfee

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

type queryHandler struct {
	*Handler
	fees  []*appfeemwpb.Fee
	infos []*npool.AppFee
	apps  map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := func() (_appIDs []string) {
		for _, fee := range h.fees {
			_appIDs = append(_appIDs, fee.AppID)
		}
		return
	}()
	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, 0, int32(len(appIDs)))
	if err != nil {
		return err
	}
	h.apps = map[string]*appmwpb.App{}
	for _, app := range apps {
		h.apps[app.EntID] = app
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, fee := range h.fees {
		app, ok := h.apps[fee.AppID]
		if !ok {
			continue
		}
		h.infos = append(h.infos, &npool.AppFee{
			ID:               fee.ID,
			EntID:            fee.EntID,
			AppID:            fee.AppID,
			AppName:          app.Name,
			GoodID:           fee.GoodID,
			GoodName:         fee.Name,
			AppGoodID:        fee.AppGoodID,
			AppGoodName:      fee.Name,
			ProductPage:      fee.ProductPage,
			Banner:           fee.Banner,
			UnitValue:        fee.UnitValue,
			MinOrderDuration: fee.MinOrderDuration,
			CreatedAt:        fee.CreatedAt,
			UpdatedAt:        fee.UpdatedAt,
		})
	}
}

func (h *Handler) GetAppFee(ctx context.Context) (*npool.AppFee, error) {
	info, err := appfeemwcli.GetFee(ctx, *h.AppGoodID)
	if err != nil {
		return nil, err
	}
	handler := &queryHandler{
		Handler: h,
		fees:    []*appfeemwpb.Fee{info},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetAppFees(ctx context.Context) ([]*npool.AppFee, uint32, error) {
	infos, total, err := appfeemwcli.GetFees(ctx, &appfeemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	handler := &queryHandler{
		Handler: h,
		fees:    infos,
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, total, nil
}
