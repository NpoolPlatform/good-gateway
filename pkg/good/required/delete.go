package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (h *Handler) DeleteRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	return requiredmwcli.DeleteRequired(ctx, *h.ID)
}
