package comment

import (
	"context"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/comment"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
	purchasedUser bool
	trialUser     bool
}

func (h *Handler) CreateComment(ctx context.Context) (*npool.Comment, error) {
	handler := &createHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUserWithUserID(ctx, *h.CommentUserID); err != nil {
		return nil, err
	}
	if err := handler.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkOrder(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	// TODO: check if trial user

	if err := commentmwcli.CreateComment(ctx, &commentmwpb.CommentReq{
		EntID:         h.EntID,
		UserID:        h.CommentUserID,
		AppGoodID:     h.AppGoodID,
		OrderID:       h.OrderID,
		Content:       h.Content,
		ReplyToID:     h.ReplyToID,
		Anonymous:     h.Anonymous,
		PurchasedUser: &handler.purchasedUser,
		TrialUser:     &handler.trialUser,
		Score:         h.Score,
	}); err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
