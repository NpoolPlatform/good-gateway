package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (h *Handler) DeleteBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	info, err := brandmwcli.GetBrandOnly(ctx, &brandmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	return brandmwcli.DeleteBrand(ctx, *h.ID)
}
