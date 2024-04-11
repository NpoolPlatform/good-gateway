package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (h *Handler) GetDeviceInfos(ctx context.Context) ([]*devicetypemwpb.DeviceInfo, uint32, error) {
	return devicetypemwcli.GetDeviceInfos(ctx, &devicetypemwpb.Conds{}, h.Offset, h.Limit)
}
