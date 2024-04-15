package powerrental

import (
	"context"

	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, err
	}
	if err := apppowerrentalmwcli.DeletePowerRental(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
