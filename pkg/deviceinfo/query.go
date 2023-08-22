package deviceinfo

import (
	"context"

	deviceinfomwcli "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
)

func (h *Handler) GetDeviceInfos(ctx context.Context) ([]*deviceinfomwpb.DeviceInfo, uint32, error) {
	return deviceinfomwcli.GetDeviceInfos(ctx, &deviceinfomwpb.Conds{}, h.Offset, h.Limit)
}
