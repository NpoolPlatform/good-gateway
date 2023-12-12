package deviceinfo

import (
	"context"

	deviceinfomwcli "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
)

func (h *Handler) DeleteDeviceInfo(ctx context.Context) (*deviceinfomwpb.DeviceInfo, error) {
	info, err := deviceinfomwcli.GetDeviceInfoOnly(ctx, &deviceinfomwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	return deviceinfomwcli.DeleteDeviceInfo(ctx, *h.ID)
}
