package comment

import (
	"context"
	"fmt"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/comment"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

func (h *Handler) UpdateComment(ctx context.Context) (*npool.Comment, error) {
	exist, err := commentmwcli.ExistCommentConds(ctx, &commentmwpb.Conds{
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

	if _, err := commentmwcli.UpdateComment(ctx, &commentmwpb.CommentReq{
		ID:      h.ID,
		Content: h.Content,
	}); err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
