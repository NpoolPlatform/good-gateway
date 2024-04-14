package score

import (
	"context"
	"fmt"

	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/score"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkScore(ctx context.Context) error {
	exist, err := scoremwcli.ExistScoreConds(ctx, &scoremwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid score")
	}
	return nil
}
