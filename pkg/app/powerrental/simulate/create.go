package simulate

import (
	"context"

	simulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"

	"github.com/google/uuid"
)

func (h *Handler) CreateSimulate(ctx context.Context) (*npool.Simulate, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := simulatemwcli.CreateSimulate(ctx, &simulatemwpb.SimulateReq{
		EntID:              h.EntID,
		AppGoodID:          h.AppGoodID,
		FixedOrderUnits:    h.FixedOrderUnits,
		FixedOrderDuration: h.FixedOrderDuration,
	}); err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
