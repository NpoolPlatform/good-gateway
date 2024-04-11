package devicetype

import (
	"context"
	"fmt"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (h *Handler) DeleteDeviceInfo(ctx context.Context) (*devicetypemwpb.DeviceInfo, error) {
	info, err := devicetypemwcli.GetDeviceInfoOnly(ctx, &devicetypemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid devicetype")
	}

	return devicetypemwcli.DeleteDeviceInfo(ctx, *h.ID)
}
