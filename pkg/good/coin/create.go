package goodcoin

import (
	"context"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodCoin(ctx context.Context) (*goodcoinmwpb.GoodCoin, error) {
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
		return nil, err
	}
	return h.GetGoodCoin(ctx)
}
