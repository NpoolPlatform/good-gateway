package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

func (h *Handler) DeleteTopMost(ctx context.Context) (*npool.TopMost, error) {
	if _, err := topmostmwcli.DeleteTopMost(ctx, *h.ID); err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
