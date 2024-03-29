package comment

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/comment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/order"

	"github.com/google/uuid"
)

func (h *Handler) CreateComment(ctx context.Context) (*npool.Comment, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	exist, err = appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid appgood")
	}

	if h.OrderID != nil {
		exist, err = ordermwcli.ExistOrderConds(ctx, &ordermwpb.Conds{
			EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID},
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
			AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
			UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		})
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, fmt.Errorf("order not matched")
		}
	}

	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if _, err := commentmwcli.CreateComment(ctx, &commentmwpb.CommentReq{
		EntID:     h.EntID,
		AppID:     h.AppID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		OrderID:   h.OrderID,
		Content:   h.Content,
		ReplyToID: h.ReplyToID,
		Anonymous: h.Anonymous,
	}); err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
