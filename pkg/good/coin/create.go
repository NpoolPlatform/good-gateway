package goodcoin

import (
	"context"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodCoins []*npool.GoodCoin
}

func (h *createHandler) getGoodCoins(ctx context.Context) error {
	h.Limit = constant.DefaultRowLimit
	for {
		goodCoins, _, err := h.GetGoodCoins(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(goodCoins) == 0 {
			return nil
		}
		h.goodCoins = append(h.goodCoins, goodCoins...)
		h.Offset += h.Limit
	}
}

func (h *createHandler) validateCandidateCoin(ctx context.Context) error {
	if len(h.goodCoins) == 0 {
		return nil
	}
	exist, err := coinmwcli.ExistCoinConds(ctx, &coinmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID},
		ENV:   &basetypes.StringVal{Op: cruder.EQ, Value: h.goodCoins[0].CoinENV},
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid coin")
	}
	return nil
}

func (h *Handler) CreateGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getGoodCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateCandidateCoin(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := goodcoinmwcli.CreateGoodCoin(ctx, &goodcoinmwpb.GoodCoinReq{
		EntID:      h.EntID,
		GoodID:     h.GoodID,
		CoinTypeID: h.CoinTypeID,
		Main:       h.Main,
		Index:      h.Index,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetGoodCoin(ctx)
}
