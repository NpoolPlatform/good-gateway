package comment

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/order"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkUser(ctx context.Context, userID string) error {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, userID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid user")
	}
	return nil
}

func (h *checkHandler) checkAppGood(ctx context.Context) error {
	exist, err := appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appgood")
	}
	return nil
}

func (h *checkHandler) checkOrder(ctx context.Context) error {
	conds := &ordermwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	}
	if h.OrderID != nil {
		conds.EntID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID}
	}
	exist, err := ordermwcli.ExistOrderConds(ctx, conds)
	if err != nil {
		return err
	}
	if !exist && h.OrderID != nil {
		return fmt.Errorf("order not matched")
	}
	return nil
}

func (h *checkHandler) checkUserComment(ctx context.Context) error {
	conds := &commentmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.TargetUserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.TargetUserID}
	} else {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	exist, err := commentmwcli.ExistCommentConds(ctx, conds)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid comment")
	}
	return nil
}
