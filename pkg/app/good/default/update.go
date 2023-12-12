package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
	defaultmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (h *Handler) UpdateDefault(ctx context.Context) (*npool.Default, error) {
	info, err := defaultmwcli.GetDefaultOnly(ctx, &defaultmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if _, err := defaultmwcli.UpdateDefault(ctx, &defaultmwpb.DefaultReq{
		ID:        h.ID,
		AppGoodID: h.AppGoodID,
	}); err != nil {
		return nil, err
	}

	return h.GetDefault(ctx)
}
