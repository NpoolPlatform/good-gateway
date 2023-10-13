package like

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"

	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/like"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/like"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"
)

func (h *Handler) UpdateLike(ctx context.Context) (*npool.Like, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	exist, err = likemwcli.ExistLikeConds(ctx, &likemwpb.Conds{
		ID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.ID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid like")
	}

	if _, err := likemwcli.UpdateLike(ctx, &likemwpb.LikeReq{
		ID:   h.ID,
		Like: h.Like,
	}); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
