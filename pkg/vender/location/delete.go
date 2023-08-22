package location

import (
	"context"
	"fmt"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (h *Handler) DeleteLocation(ctx context.Context) (*locationmwpb.Location, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	return locationmwcli.DeleteLocation(ctx, *h.ID)
}
