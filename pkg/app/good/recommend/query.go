package recommend

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

type queryHandler struct {
	*Handler
	recommends []*recommendmwpb.Recommend
	infos      []*npool.Recommend
	apps       map[string]*appmwpb.App
	users      map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, recommend := range h.recommends {
		appIDs = append(appIDs, recommend.AppID)
	}
	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, int32(0), int32(len(appIDs)))
	if err != nil {
		return err
	}
	for _, app := range apps {
		h.apps[app.EntID] = app
	}
	return nil
}

func (h *queryHandler) getUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, recommend := range h.recommends {
		userIDs = append(userIDs, recommend.RecommenderID)
	}
	users, _, err := usermwcli.GetUsers(ctx, &usermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}, int32(0), int32(len(userIDs)))
	if err != nil {
		return err
	}
	for _, user := range users {
		h.users[user.EntID] = user
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, recommend := range h.recommends {
		info := &npool.Recommend{
			ID:             recommend.ID,
			EntID:          recommend.EntID,
			AppID:          recommend.AppID,
			RecommenderID:  recommend.RecommenderID,
			GoodID:         recommend.GoodID,
			GoodName:       recommend.GoodName,
			RecommendIndex: recommend.RecommendIndex,
			Message:        recommend.Message,
			CreatedAt:      recommend.CreatedAt,
			UpdatedAt:      recommend.UpdatedAt,
		}

		app, ok := h.apps[recommend.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[recommend.RecommenderID]
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

func (h *Handler) GetRecommend(ctx context.Context) (*npool.Recommend, error) {
	recommend, err := recommendmwcli.GetRecommend(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if recommend == nil {
		return nil, fmt.Errorf("invalid recommend")
	}

	handler := &queryHandler{
		Handler:    h,
		recommends: []*recommendmwpb.Recommend{recommend},
		apps:       map[string]*appmwpb.App{},
		users:      map[string]*usermwpb.User{},
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

func (h *Handler) GetRecommends(ctx context.Context) ([]*npool.Recommend, uint32, error) {
	conds := &recommendmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.RecommenderID != nil {
		conds.RecommenderID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID}
	}

	recommends, total, err := recommendmwcli.GetRecommends(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(recommends) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:    h,
		recommends: recommends,
		apps:       map[string]*appmwpb.App{},
		users:      map[string]*usermwpb.User{},
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
