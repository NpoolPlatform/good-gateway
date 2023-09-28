package score

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/score"
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

	id := uuid.NewString()
	if h.ID == nil {
		h.ID = &id
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
