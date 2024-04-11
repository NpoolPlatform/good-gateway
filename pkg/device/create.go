package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"

	"github.com/google/uuid"
)

func (h *Handler) CreateDeviceInfo(ctx context.Context) (*devicetypemwpb.DeviceInfo, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	return devicetypemwcli.CreateDeviceInfo(ctx, &devicetypemwpb.DeviceInfoReq{
		EntID:            h.EntID,
		Type:             h.Type,
		Manufacturer:     h.Manufacturer,
		PowerConsumption: h.PowerConsumption,
		ShipmentAt:       h.ShipmentAt,
		Posters:          h.Posters,
	})
}
