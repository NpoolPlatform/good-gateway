package simulate

import (
	"context"
	"fmt"

	simulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

func (h *Handler) UpdateSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := simulatemwcli.GetSimulateOnly(ctx, &simulatemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid simulate")
	}

	if _, err := simulatemwcli.UpdateSimulate(ctx, &simulatemwpb.SimulateReq{
		ID:                 h.ID,
		FixedOrderUnits:    h.FixedOrderUnits,
		FixedOrderDuration: h.FixedOrderDuration,
	}); err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
