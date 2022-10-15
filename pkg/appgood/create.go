package appgood

import (
	"context"

	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"
)

func CreateAppGood(
	ctx context.Context,
	appID, goodID string, online, visible bool,
	goodName string, price decimal.Decimal,
	displayIndex, purchaseLimit, commissionPercent int32,
) (*npool.Good, error) {
	return nil, nil
}
