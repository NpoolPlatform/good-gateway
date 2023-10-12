package like

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/like"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/like"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"
)

type queryHandler struct {
	*Handler
	likes []*likemwpb.Like
	infos []*npool.Like
	apps  map[string]*appmwpb.App
	users map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, like := range h.likes {
		appIDs = append(appIDs, like.AppID)
	}
	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, int32(0), int32(len(appIDs)))
	if err != nil {
		return err
	}
	for _, app := range apps {
		h.apps[app.ID] = app
	}
	return nil
}

func (h *queryHandler) getUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, like := range h.likes {
		userIDs = append(userIDs, like.UserID)
	}
	users, _, err := usermwcli.GetUsers(ctx, &usermwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}, int32(0), int32(len(userIDs)))
	if err != nil {
		return err
	}
	for _, user := range users {
		h.users[user.ID] = user
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, like := range h.likes {
		info := &npool.Like{
			ID:        like.ID,
			AppID:     like.AppID,
			UserID:    like.UserID,
			AppGoodID: like.AppGoodID,
			GoodName:  like.GoodName,
			Like:      like.Like,
			CreatedAt: like.CreatedAt,
		}

		app, ok := h.apps[like.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[like.UserID]
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

func (h *Handler) GetLike(ctx context.Context) (*npool.Like, error) {
	like, err := likemwcli.GetLike(ctx, *h.ID)
	if err != nil {
		return nil, err
	}
	if like == nil {
		return nil, fmt.Errorf("invalid like")
	}

	handler := &queryHandler{
		Handler: h,
		likes:   []*likemwpb.Like{like},
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
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

func (h *Handler) GetLikes(ctx context.Context) ([]*npool.Like, uint32, error) {
	if h.UserID != nil {
		exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
		if err != nil {
			return nil, 0, err
		}
		if !exist {
			return nil, 0, fmt.Errorf("invalid user")
		}
	}

	conds := &likemwpb.Conds{}
	conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	likes, total, err := likemwcli.GetLikes(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(likes) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		likes:   likes,
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
