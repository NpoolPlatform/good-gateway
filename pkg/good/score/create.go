package score

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appgooodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/score"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/score"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateScore(ctx context.Context) (*npool.Score, error) {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	exist, err = appgooodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid appgood")
	}

	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	if h.Score != nil {
		maxScore := decimal.RequireFromString("10.0")
		score, err := decimal.NewFromString(*h.Score)
		if err != nil {
			return nil, err
		}
		if score.Cmp(maxScore) > 0 {
			return nil, fmt.Errorf("invalid score")
		}
	}

	if _, err := scoremwcli.CreateScore(ctx, &scoremwpb.ScoreReq{
		ID:        h.ID,
		AppID:     h.AppID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Score:     h.Score,
	}); err != nil {
		return nil, err
	}

	return h.GetScore(ctx)
}
