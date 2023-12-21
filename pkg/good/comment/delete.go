package comment

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/comment"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) checkUser(ctx context.Context) error {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.TargetUserID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid user")
	}
	if h.UserID != nil {
		exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid user")
		}
	}
	return nil
}

func (h *deleteHandler) checkComment(ctx context.Context) error {
	exist, err := commentmwcli.ExistCommentConds(ctx, &commentmwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.TargetUserID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid comment")
	}
	return nil
}

func (h *Handler) DeleteComment(ctx context.Context) (*npool.Comment, error) {
	handler := &deleteHandler{
		Handler: h,
	}
	if err := handler.checkUser(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkComment(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetComment(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := commentmwcli.DeleteComment(ctx, *h.ID); err != nil {
		return nil, err
	}
	return info, nil
}
