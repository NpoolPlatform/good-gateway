package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (h *Handler) UpdateLocation(ctx context.Context) (*locationmwpb.Location, error) {
	info, err := locationmwcli.GetLocationOnly(ctx, &locationmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	return locationmwcli.UpdateLocation(ctx, &locationmwpb.LocationReq{
		ID:       h.ID,
		Country:  h.Country,
		Province: h.Province,
		City:     h.City,
		Address:  h.Address,
		BrandID:  h.BrandID,
	})
}
