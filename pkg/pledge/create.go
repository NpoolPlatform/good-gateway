package pledge

import (
	"context"
	"fmt"

	contractmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/contract"
	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	pledgemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/pledge"
	contractmwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"
	accounttypes "github.com/NpoolPlatform/message/npool/basetypes/account/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/pledge"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
	pledgemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodCoinName               *string
	contractDevelopmentAddress *string
	contractCalculateAddress   *string
}

func (h *createHandler) getCoin(ctx context.Context) error {
	coin, err := coinmwcli.GetCoin(ctx, *h.CoinTypeID)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}

	h.goodCoinName = &coin.Name

	return nil
}

func (h *createHandler) createGoodCoin(ctx context.Context) error {
	main := true
	if err := goodcoinmwcli.CreateGoodCoin(ctx, &goodcoinmwpb.GoodCoinReq{
		GoodID:     h.GoodID,
		CoinTypeID: h.CoinTypeID,
		Main:       &main,
	}); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createDevelopmentAddress(ctx context.Context) error {
	acc, err := sphinxproxycli.CreateAddress(ctx, *h.goodCoinName)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("fail create address")
	}
	h.contractDevelopmentAddress = &acc.Address
	_, err = contractmwcli.CreateAccount(ctx, &contractmwpb.AccountReq{
		GoodID:       h.GoodID,
		PledgeID:     h.EntID,
		ContractType: accounttypes.ContractType_ContractDeployment.Enum(),
		CoinTypeID:   h.CoinTypeID,
		Address:      h.contractDevelopmentAddress,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createCalculateAddress(ctx context.Context) error {
	acc, err := sphinxproxycli.CreateAddress(ctx, *h.goodCoinName)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("fail create address")
	}
	h.contractCalculateAddress = &acc.Address
	_, err = contractmwcli.CreateAccount(ctx, &contractmwpb.AccountReq{
		GoodID:       h.GoodID,
		PledgeID:     h.EntID,
		ContractType: accounttypes.ContractType_ContractCalculate.Enum(),
		CoinTypeID:   h.CoinTypeID,
		Address:      h.contractDevelopmentAddress,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreatePledge(ctx context.Context) (*npool.Pledge, error) {
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoin(ctx); err != nil {
		return nil, err
	}
	if err := handler.createDevelopmentAddress(ctx); err != nil {
		return nil, err
	}
	if err := handler.createCalculateAddress(ctx); err != nil {
		return nil, err
	}
	if err := pledgemwcli.CreatePledge(ctx, &pledgemwpb.PledgeReq{
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		GoodType:             h.GoodType,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		ContractCodeURL:      h.ContractCodeURL,
		ContractCodeBranch:   h.ContractCodeBranch,
		ContractState:        h.ContractState,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.createGoodCoin(ctx); err != nil {
		return nil, err
	}
	return h.GetPledge(ctx)
}
