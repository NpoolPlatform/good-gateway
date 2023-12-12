package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

	"github.com/google/uuid"
)

func (h *Handler) CreateLocation(ctx context.Context) (*locationmwpb.Location, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	return locationmwcli.CreateLocation(ctx, &locationmwpb.LocationReq{
		EntID:    h.EntID,
		Country:  h.Country,
		Province: h.Province,
		City:     h.City,
		Address:  h.Address,
		BrandID:  h.BrandID,
	})
}
