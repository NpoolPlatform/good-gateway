package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (h *Handler) UpdateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	info, err := requiredmwcli.GetRequiredOnly(ctx, &requiredmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	return requiredmwcli.UpdateRequired(ctx, &requiredmwpb.RequiredReq{
		ID:   h.ID,
		Must: h.Must,
	})
}
