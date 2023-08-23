package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"

	"github.com/google/uuid"
)

func (h *Handler) CreateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
	}

	return requiredmwcli.CreateRequired(ctx, &requiredmwpb.RequiredReq{
		ID:             h.ID,
		MainGoodID:     h.MainGoodID,
		RequiredGoodID: h.RequiredGoodID,
		Must:           h.Must,
	})
}
