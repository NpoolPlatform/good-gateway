package score

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/score"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
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

	if _, err := scoremwcli.CreateScore(ctx, &scoremwpb.ScoreReq{
		ID:     h.ID,
		AppID:  h.AppID,
		UserID: h.UserID,
		GoodID: h.GoodID,
		Score:  h.Score,
	}); err != nil {
		return nil, err
	}

	return h.GetScore(ctx)
}
