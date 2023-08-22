package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (h *Handler) GetLocations(ctx context.Context) ([]*locationmwpb.Location, uint32, error) {
	return locationmwcli.GetLocations(ctx, &locationmwpb.Conds{}, h.Offset, h.Limit)
}
