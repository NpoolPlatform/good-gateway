package deviceinfo

import (
	"context"
	"fmt"

	deviceinfomwcli "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
)

func (h *Handler) UpdateDeviceInfo(ctx context.Context) (*deviceinfomwpb.DeviceInfo, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	return deviceinfomwcli.UpdateDeviceInfo(ctx, &deviceinfomwpb.DeviceInfoReq{
		ID:               h.ID,
		Type:             h.Type,
		Manufacturer:     h.Manufacturer,
		PowerConsumption: h.PowerConsumption,
		ShipmentAt:       h.ShipmentAt,
		Posters:          h.Posters,
	})
}
