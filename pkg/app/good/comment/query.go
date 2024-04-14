package comment

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/comment"
	commentmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"

	"github.com/google/uuid"
)

type queryHandler struct {
	*checkHandler
	comments []*commentmwpb.Comment
	infos    []*npool.Comment
	apps     map[string]*appmwpb.App
	users    map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, comment := range h.comments {
			appIDs = append(appIDs, comment.AppID)
		}
		return
	}())
	return nil
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = goodgwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, comment := range h.comments {
			userIDs = append(userIDs, comment.UserID)
		}
		return
	}())
	return nil
}

func (h *queryHandler) formalize() {
	for _, comment := range h.comments {
		info := &npool.Comment{
			ID:            comment.ID,
			EntID:         comment.EntID,
			AppID:         comment.AppID,
			UserID:        comment.UserID,
			GoodID:        comment.GoodID,
			AppGoodID:     comment.AppGoodID,
			GoodName:      comment.GoodName,
			Content:       comment.Content,
			Anonymous:     comment.Anonymous,
			PurchasedUser: comment.PurchasedUser,
			TrialUser:     comment.TrialUser,
			Score:         comment.Score,
			Hide:          comment.Hide,
			HideReason:    comment.HideReason,
			CreatedAt:     comment.CreatedAt,
			UpdatedAt:     comment.UpdatedAt,
		}

		if _, err := uuid.Parse(comment.OrderID); err == nil {
			if comment.OrderID != uuid.Nil.String() {
				info.OrderID = &comment.OrderID
			}
		}
		if _, err := uuid.Parse(comment.ReplyToID); err == nil {
			if comment.ReplyToID != uuid.Nil.String() {
				info.ReplyToID = &comment.ReplyToID
			}
		}

		app, ok := h.apps[comment.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[comment.UserID]
		if ok {
			if user.Username != "" {
				info.Username = &user.Username
			}
			if user.EmailAddress != "" {
				info.EmailAddress = &user.EmailAddress
			}
			if user.PhoneNO != "" {
				info.PhoneNO = &user.PhoneNO
			}
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetComment(ctx context.Context) (*npool.Comment, error) {
	comment, err := commentmwcli.GetComment(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("invalid comment")
	}

	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		comments: []*commentmwpb.Comment{comment},
		apps:     map[string]*appmwpb.App{},
		users:    map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetComments(ctx context.Context) ([]*npool.Comment, uint32, error) {
	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		apps:  map[string]*appmwpb.App{},
		users: map[string]*usermwpb.User{},
	}
	if h.UserID != nil {
		if err := handler.CheckUser(ctx); err != nil {
			return nil, 0, fmt.Errorf("invalid user")
		}
	}

	conds := &commentmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	comments, total, err := commentmwcli.GetComments(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(comments) == 0 {
		return nil, total, nil
	}

	handler.comments = comments
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
