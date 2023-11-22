package topmost

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

type queryHandler struct {
	*Handler
	topmosts []*topmostmwpb.TopMost
	infos    []*npool.TopMost
	apps     map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, topmost := range h.topmosts {
		appIDs = append(appIDs, topmost.AppID)
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

func (h *queryHandler) formalize() {
	for _, topmost := range h.topmosts {
		info := &npool.TopMost{
			ID:                     topmost.ID,
			AppID:                  topmost.AppID,
			TopMostType:            topmost.TopMostType,
			Title:                  topmost.Title,
			Message:                topmost.Message,
			Posters:                topmost.Posters,
			StartAt:                topmost.StartAt,
			EndAt:                  topmost.EndAt,
			ThresholdCredits:       topmost.ThresholdCredits,
			RegisterElapsedSeconds: topmost.RegisterElapsedSeconds,
			ThresholdPurchases:     topmost.ThresholdPurchases,
			ThresholdPaymentAmount: topmost.ThresholdPaymentAmount,
			KycMust:                topmost.KycMust,
			CreatedAt:              topmost.CreatedAt,
			UpdatedAt:              topmost.UpdatedAt,
		}

		app, ok := h.apps[topmost.AppID]
		if ok {
			info.AppName = app.Name
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetTopMost(ctx context.Context) (*npool.TopMost, error) {
	info, err := topmostmwcli.GetTopMost(ctx, *h.ID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &queryHandler{
		Handler:  h,
		topmosts: []*topmostmwpb.TopMost{info},
		apps:     map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetTopMosts(ctx context.Context) ([]*npool.TopMost, uint32, error) {
	infos, total, err := topmostmwcli.GetTopMosts(
		ctx,
		&topmostmwpb.Conds{
			AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:  h,
		topmosts: infos,
		apps:     map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
