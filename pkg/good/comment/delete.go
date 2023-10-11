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

func (h *Handler) DeleteComment(ctx context.Context) (*npool.Comment, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	exist, err = commentmwcli.ExistCommentConds(ctx, &commentmwpb.Conds{
		ID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.ID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid comment")
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
