package like

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/like"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/like"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

func (h *Handler) CreateLike(ctx context.Context) (*npool.Like, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
	}

	if _, err := likemwcli.CreateLike(ctx, &likemwpb.LikeReq{
		ID:        h.ID,
		AppID:     h.AppID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Like:      h.Like,
	}); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
