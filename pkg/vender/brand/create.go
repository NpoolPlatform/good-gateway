package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"

	"github.com/google/uuid"
)

func (h *Handler) CreateBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
	}

	return brandmwcli.CreateBrand(ctx, &brandmwpb.BrandReq{
		ID:   h.ID,
		Name: h.Name,
		Logo: h.Logo,
	})
}
