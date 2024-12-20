package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	apppledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/pledge"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"
	goodcoingwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin/reward"
	apppledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

type queryHandler struct {
	*Handler
	appPledges []*apppledgemwpb.Pledge
	infos      []*npool.AppPledge
	apps       map[string]*appmwpb.App
	appCoins   map[string]map[string]*appcoinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appPledge := range h.appPledges {
			appIDs = append(appIDs, appPledge.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.appCoins = map[string]map[string]*appcoinmwpb.Coin{}
	appCoinTypeIDs := map[string][]string{}
	for _, appPledge := range h.appPledges {
		coinTypeIDs, ok := appCoinTypeIDs[appPledge.AppID]
		if !ok {
			coinTypeIDs = []string{}
		}
		for _, goodCoin := range appPledge.GoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		appCoinTypeIDs[appPledge.AppID] = coinTypeIDs
	}
	for appID, coinTypeIDs := range appCoinTypeIDs {
		coins, err := goodgwcommon.GetAppCoins(ctx, appID, coinTypeIDs)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.appCoins[appID] = coins
	}
	return wlog.WrapError(err)
}

//nolint:funlen
func (h *queryHandler) formalize() {
	for _, appPledge := range h.appPledges {
		info := &npool.AppPledge{
			ID:        appPledge.ID,
			EntID:     appPledge.EntID,
			AppID:     appPledge.AppID,
			GoodID:    appPledge.GoodID,
			AppGoodID: appPledge.AppGoodID,

			GoodType:             appPledge.GoodType,
			BenefitType:          appPledge.BenefitType,
			GoodName:             appPledge.GoodName,
			ServiceStartAt:       appPledge.AppGoodServiceStartAt,
			GoodStartMode:        appPledge.GoodStartMode,
			TestOnly:             appPledge.TestOnly,
			BenefitIntervalHours: appPledge.BenefitIntervalHours,
			GoodPurchasable:      appPledge.GoodPurchasable,
			GoodOnline:           appPledge.GoodOnline,
			State:                appPledge.State,

			AppGoodPurchasable:  appPledge.AppGoodPurchasable,
			AppGoodOnline:       appPledge.AppGoodOnline,
			EnableProductPage:   appPledge.EnableProductPage,
			ProductPage:         appPledge.ProductPage,
			Visible:             appPledge.Visible,
			AppGoodName:         appPledge.AppGoodName,
			DisplayIndex:        appPledge.DisplayIndex,
			Banner:              appPledge.Banner,
			EnableSetCommission: appPledge.EnableSetCommission,
			AppGoodStartMode:    appPledge.AppGoodStartMode,

			Likes:          appPledge.Likes,
			Dislikes:       appPledge.Dislikes,
			Score:          appPledge.Score,
			ScoreCount:     appPledge.ScoreCount,
			RecommendCount: appPledge.RecommendCount,
			CommentCount:   appPledge.CommentCount,

			LastRewardAt: appPledge.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range appPledge.Rewards {
					coins, ok := h.appCoins[appPledge.AppID]
					if !ok {
						continue
					}
					coin, ok := coins[reward.CoinTypeID]
					if !ok {
						continue
					}
					rewards = append(rewards, &goodcoinrewardgwpb.RewardInfo{
						CoinTypeID:            reward.CoinTypeID,
						CoinName:              coin.Name,
						CoinUnit:              coin.Unit,
						CoinENV:               coin.ENV,
						CoinLogo:              coin.Logo,
						RewardTID:             reward.RewardTID,
						NextRewardStartAmount: reward.NextRewardStartAmount,
						LastRewardAmount:      reward.LastRewardAmount,
						LastUnitRewardAmount:  reward.LastUnitRewardAmount,
						TotalRewardAmount:     reward.TotalRewardAmount,
						MainCoin:              reward.MainCoin,
					})
				}
				return
			}(),

			GoodCoins: func() (coins []*goodcoingwpb.GoodCoinInfo) {
				for _, goodCoin := range appPledge.GoodCoins {
					appCoins, ok := h.appCoins[appPledge.AppID]
					if !ok {
						continue
					}
					appCoin, ok := appCoins[goodCoin.CoinTypeID]
					if !ok {
						continue
					}
					coins = append(coins, &goodcoingwpb.GoodCoinInfo{
						CoinTypeID: goodCoin.CoinTypeID,
						CoinName:   appCoin.Name,
						CoinUnit:   appCoin.Unit,
						CoinENV:    appCoin.ENV,
						CoinLogo:   appCoin.Logo,
						Main:       goodCoin.Main,
						Index:      goodCoin.Index,
					})
				}
				return
			}(),
			Descriptions:  appPledge.Descriptions,
			Posters:       appPledge.Posters,
			DisplayNames:  appPledge.DisplayNames,
			DisplayColors: appPledge.DisplayColors,
			Labels:        appPledge.Labels,

			CreatedAt: appPledge.CreatedAt,
			UpdatedAt: appPledge.UpdatedAt,
		}

		app, ok := h.apps[appPledge.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPledge(ctx context.Context) (*npool.AppPledge, error) {
	appPledge, err := apppledgemwcli.GetPledge(ctx, *h.AppGoodID)
	if err != nil {
		return nil, err
	}
	if appPledge == nil {
		return nil, wlog.Errorf("invalid apppledge")
	}

	handler := &queryHandler{
		Handler:    h,
		appPledges: []*apppledgemwpb.Pledge{appPledge},
		apps:       map[string]*appmwpb.App{},
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

func (h *Handler) GetPledges(ctx context.Context) ([]*npool.AppPledge, uint32, error) {
	appPledges, total, err := apppledgemwcli.GetPledges(ctx, &apppledgemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(appPledges) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:    h,
		appPledges: appPledges,
		apps:       map[string]*appmwpb.App{},
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
