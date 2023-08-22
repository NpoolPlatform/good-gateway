package brand

import (
	"context"
	"fmt"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (h *Handler) DeleteBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	return brandmwcli.DeleteBrand(ctx, *h.ID)
}
