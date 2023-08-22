package deviceinfo

import (
	"context"

	deviceinfomwcli "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"

	"github.com/google/uuid"
)

func (h *Handler) CreateDeviceInfo(ctx context.Context) (*deviceinfomwpb.DeviceInfo, error) {
	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
	}

	return deviceinfomwcli.CreateDeviceInfo(ctx, &deviceinfomwpb.DeviceInfoReq{
		ID:               h.ID,
		Type:             h.Type,
		Manufacturer:     h.Manufacturer,
		PowerConsumption: h.PowerConsumption,
		ShipmentAt:       h.ShipmentAt,
		Posters:          h.Posters,
	})
}
