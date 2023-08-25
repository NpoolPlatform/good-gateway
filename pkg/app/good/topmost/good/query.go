package topmostgood

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

type queryHandler struct {
	*Handler
	goods []*topmostgoodmwpb.TopMostGood
	infos []*npool.TopMostGood
	apps  map[string]*appmwpb.App
	coins map[string]*coinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, good := range h.goods {
		appIDs = append(appIDs, good.AppID)
	}
	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, int32(0), int32(len(appIDs)))
	if err != nil {
		return err
	}
	for _, app := range apps {
		h.apps[app.ID] = app
	}
	return nil
}

func (h *queryHandler) getCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, good := range h.goods {
		coinTypeIDs = append(coinTypeIDs, good.CoinTypeID)
	}
	coins, _, err := coinmwcli.GetCoins(ctx, &coinmwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}, int32(0), int32(len(coinTypeIDs)))
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.coins[coin.ID] = coin
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, good := range h.goods {
		info := &npool.TopMostGood{
			ID:             good.ID,
			AppID:          good.AppID,
			GoodID:         good.GoodID,
			GoodName:       good.GoodName,
			AppGoodID:      good.AppGoodID,
			AppGoodName:    good.AppGoodName,
			CoinTypeID:     good.CoinTypeID,
			TopMostID:      good.TopMostID,
			TopMostType:    good.TopMostType,
			TopMostTitle:   good.TopMostTitle,
			TopMostMessage: good.TopMostMessage,
			Price:          good.Price,
			Posters:        good.Posters,
			CreatedAt:      good.CreatedAt,
			UpdatedAt:      good.UpdatedAt,
		}

		app, ok := h.apps[good.AppID]
		if ok {
			info.AppName = app.Name
		}
		coin, ok := h.coins[good.CoinTypeID]
		if ok {
			info.CoinName = coin.Name
			info.CoinLogo = coin.Logo
			info.CoinEnv = coin.ENV
			info.CoinUnit = coin.Unit
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	info, err := topmostgoodmwcli.GetTopMostGood(ctx, *h.ID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &queryHandler{
		Handler: h,
		goods:   []*topmostgoodmwpb.TopMostGood{info},
		apps:    map[string]*appmwpb.App{},
		coins:   map[string]*coinmwpb.Coin{},
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

func (h *Handler) GetTopMostGoods(ctx context.Context) ([]*npool.TopMostGood, uint32, error) {
	infos, total, err := topmostgoodmwcli.GetTopMostGoods(
		ctx,
		&topmostgoodmwpb.Conds{
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
		Handler: h,
		goods:   infos,
		apps:    map[string]*appmwpb.App{},
		coins:   map[string]*coinmwpb.Coin{},
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
