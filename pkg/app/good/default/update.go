package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/default"
	defaultmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (h *Handler) UpdateDefault(ctx context.Context) (*npool.Default, error) {
	if _, err := defaultmwcli.UpdateDefault(ctx, &defaultmwpb.DefaultReq{
		ID:        h.ID,
		AppGoodID: h.AppGoodID,
	}); err != nil {
		return nil, err
	}

	return h.GetDefault(ctx)
}
