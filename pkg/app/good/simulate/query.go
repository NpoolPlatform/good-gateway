package simulate

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	simulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/simulate"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/simulate"
	simulatemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
)

type queryHandler struct {
	*Handler
	simulates []*simulatemwpb.Simulate
	infos     []*npool.Simulate
	apps      map[string]*appmwpb.App
	coins     map[string]*coinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, def := range h.simulates {
		appIDs = append(appIDs, def.AppID)
	}
	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, int32(0), int32(len(appIDs)))
	if err != nil {
		return err
	}
	for _, app := range apps {
		h.apps[app.EntID] = app
	}
	return nil
}

func (h *queryHandler) getCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, def := range h.simulates {
		coinTypeIDs = append(coinTypeIDs, def.CoinTypeID)
	}
	coins, _, err := coinmwcli.GetCoins(ctx, &coinmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}, int32(0), int32(len(coinTypeIDs)))
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.coins[coin.EntID] = coin
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, def := range h.simulates {
		info := &npool.Simulate{
			ID:                 def.ID,
			EntID:              def.EntID,
			AppID:              def.AppID,
			GoodID:             def.GoodID,
			GoodName:           def.GoodName,
			AppGoodID:          def.AppGoodID,
			AppGoodName:        def.AppGoodName,
			CoinTypeID:         def.CoinTypeID,
			FixedOrderUnits:    def.FixedOrderUnits,
			FixedOrderDuration: def.FixedOrderDuration,
			CreatedAt:          def.CreatedAt,
			UpdatedAt:          def.UpdatedAt,
		}

		app, ok := h.apps[def.AppID]
		if ok {
			info.AppName = app.Name
		}
		coin, ok := h.coins[def.CoinTypeID]
		if ok {
			info.CoinName = coin.Name
			info.CoinLogo = coin.Logo
			info.CoinEnv = coin.ENV
			info.CoinUnit = coin.Unit
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := simulatemwcli.GetSimulate(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid simulate")
	}

	handler := &queryHandler{
		Handler:   h,
		simulates: []*simulatemwpb.Simulate{info},
		apps:      map[string]*appmwpb.App{},
		coins:     map[string]*coinmwpb.Coin{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetSimulateExt(ctx context.Context, info *simulatemwpb.Simulate) (*npool.Simulate, error) {
	if info == nil {
		return nil, nil
	}

	handler := &queryHandler{
		Handler:   h,
		simulates: []*simulatemwpb.Simulate{info},
		apps:      map[string]*appmwpb.App{},
		coins:     map[string]*coinmwpb.Coin{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetSimulates(ctx context.Context) ([]*npool.Simulate, uint32, error) {
	infos, total, err := simulatemwcli.GetSimulates(
		ctx,
		&simulatemwpb.Conds{
			AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:   h,
		simulates: infos,
		apps:      map[string]*appmwpb.App{},
		coins:     map[string]*coinmwpb.Coin{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
