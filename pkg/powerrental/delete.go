package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"
)

func (h *Handler) DeletePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	info, err := h.GetPowerRental(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := powerrentalmwcli.DeletePowerRental(ctx, h.ID, h.EntID, h.GoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
