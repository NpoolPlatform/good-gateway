package goodcoin

import (
	"context"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (h *Handler) GetGoodCoin(ctx context.Context) (*goodcoinmwpb.GoodCoin, error) {
	return goodcoinmwcli.GetGoodCoin(ctx, *h.EntID)
}

func (h *Handler) GetGoodCoins(ctx context.Context) ([]*goodcoinmwpb.GoodCoin, uint32, error) {
	return goodcoinmwcli.GetGoodCoins(ctx, &goodcoinmwpb.Conds{}, h.Offset, h.Limit)
}
