package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
)

func (h *Handler) DeleteDefault(ctx context.Context) (*npool.Default, error) {
	info, err := h.GetDefault(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := defaultmwcli.DeleteDefault(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
