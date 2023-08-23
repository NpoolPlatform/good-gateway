package score

import (
	"context"
	"fmt"

	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/score"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/score"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"
)

func (h *Handler) DeleteScore(ctx context.Context) (*npool.Score, error) {
	exist, err := scoremwcli.ExistScoreConds(ctx, &scoremwpb.Conds{
		ID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.ID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid score")
	}

	info, err := h.GetScore(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := scoremwcli.DeleteScore(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
