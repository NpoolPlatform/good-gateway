package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (h *Handler) UpdateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	return requiredmwcli.UpdateRequired(ctx, &requiredmwpb.RequiredReq{
		ID:   h.ID,
		Must: h.Must,
	})
}
