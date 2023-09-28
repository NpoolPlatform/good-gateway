package comment

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/comment"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/comment"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"

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

	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
	}

	if _, err := commentmwcli.CreateComment(ctx, &commentmwpb.CommentReq{
		ID:        h.ID,
		AppID:     h.AppID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		OrderID:   h.OrderID,
		Content:   h.Content,
		ReplyToID: h.ReplyToID,
	}); err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
