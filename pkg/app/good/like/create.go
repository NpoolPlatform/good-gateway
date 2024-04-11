package like

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appgooodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/like"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/like"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"

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

	exist, err = appgooodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid appgood")
	}

	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := likemwcli.CreateLike(ctx, &likemwpb.LikeReq{
		EntID:     h.EntID,
		AppID:     h.AppID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Like:      h.Like,
	}); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
