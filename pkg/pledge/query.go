package pledge

import (
	"context"
	"fmt"

	contractmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/contract"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	contractmwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"
	accounttypes "github.com/NpoolPlatform/message/npool/basetypes/account/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	goodcoingwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
	pledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
)

type queryHandler struct {
	*Handler
	pledges           []*pledgemwpb.Pledge
	coins             map[string]*coinmwpb.Coin
	deploymentAddress map[string]*contractmwpb.Account
	calculateAddress  map[string]*contractmwpb.Account
	poolGoodUsers     map[string]*goodusermwpb.GoodUser
	infos             []*npool.Pledge
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = goodgwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, pledge := range h.pledges {
			for _, goodCoin := range pledge.GoodCoins {
				coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getPledgeAddress(ctx context.Context) (err error) {
	pledgeIDs := []string{}
	for _, pledge := range h.pledges {
		pledgeIDs = append(pledgeIDs, pledge.EntID)
	}
	accounts, _, err := contractmwcli.GetAccounts(ctx, &contractmwpb.Conds{
		PledgeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: pledgeIDs},
	}, int32(0), int32(len(pledgeIDs)))
	if accounts == nil {
		return nil
	}
	for _, accont := range accounts {
		if accont.ContractType == accounttypes.ContractType_ContractDeployment {
			h.deploymentAddress[accont.PledgeID] = accont
		}
		if accont.ContractType == accounttypes.ContractType_ContractCalculate {
			h.calculateAddress[accont.PledgeID] = accont
		}
	}
	return err
}

func (h *queryHandler) formalize() {
	for _, pledge := range h.pledges {
		deploymentAddress := ""
		deploymentAccount, ok := h.deploymentAddress[pledge.EntID]
		if ok {
			deploymentAddress = deploymentAccount.Address
		}
		calculateAddress := ""
		calculateAccount, ok := h.calculateAddress[pledge.EntID]
		if ok {
			calculateAddress = calculateAccount.Address
		}
		info := &npool.Pledge{
			ID:                        pledge.ID,
			EntID:                     pledge.EntID,
			GoodID:                    pledge.GoodID,
			ContractDeploymentAddress: deploymentAddress,
			ContractCalculateAddress:  calculateAddress,

			GoodType:             pledge.GoodType,
			BenefitType:          pledge.BenefitType,
			Name:                 pledge.Name,
			ServiceStartAt:       pledge.ServiceStartAt,
			StartMode:            pledge.StartMode,
			TestOnly:             pledge.TestOnly,
			BenefitIntervalHours: pledge.BenefitIntervalHours,
			Purchasable:          pledge.Purchasable,
			Online:               pledge.Online,
			State:                pledge.State,

			ContractCodeURL:    pledge.ContractCodeURL,
			ContractCodeBranch: pledge.ContractCodeBranch,
			ContractState:      pledge.ContractState,

			RewardState:  pledge.RewardState,
			LastRewardAt: pledge.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range pledge.Rewards {
					coin, ok := h.coins[reward.CoinTypeID]
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
				for _, goodCoin := range pledge.GoodCoins {
					coin, ok := h.coins[goodCoin.CoinTypeID]
					if !ok {
						continue
					}
					coins = append(coins, &goodcoingwpb.GoodCoinInfo{
						CoinTypeID: goodCoin.CoinTypeID,
						CoinName:   coin.Name,
						CoinUnit:   coin.Unit,
						CoinENV:    coin.ENV,
						CoinLogo:   coin.Logo,
						Main:       goodCoin.Main,
						Index:      goodCoin.Index,
					})
				}
				return
			}(),

			CreatedAt: pledge.CreatedAt,
			UpdatedAt: pledge.UpdatedAt,
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPledge(ctx context.Context) (*npool.Pledge, error) {
	pledge, err := pledgemwcli.GetPledge(ctx, *h.GoodID)
	if err != nil {
		return nil, err
	}
	if pledge == nil {
		return nil, fmt.Errorf("invalid pledge")
	}

	handler := &queryHandler{
		Handler:           h,
		pledges:           []*pledgemwpb.Pledge{pledge},
		coins:             map[string]*coinmwpb.Coin{},
		deploymentAddress: map[string]*contractmwpb.Account{},
		calculateAddress:  map[string]*contractmwpb.Account{},
		poolGoodUsers:     map[string]*goodusermwpb.GoodUser{},
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	if err := handler.getPledgeAddress(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPledges(ctx context.Context) ([]*npool.Pledge, uint32, error) {
	pledges, total, err := pledgemwcli.GetPledges(ctx, &pledgemwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(pledges) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:           h,
		pledges:           pledges,
		coins:             map[string]*coinmwpb.Coin{},
		deploymentAddress: map[string]*contractmwpb.Account{},
		calculateAddress:  map[string]*contractmwpb.Account{},
		poolGoodUsers:     map[string]*goodusermwpb.GoodUser{},
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	if err := handler.getPledgeAddress(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
