package goodcoin

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
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
