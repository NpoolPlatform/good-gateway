package delegatedstaking

import (
	"context"
	"fmt"

	appdelegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/delegatedstaking"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/delegatedstaking"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDelegatedStaking(ctx context.Context) error {
	exist, err := appdelegatedstakingmwcli.ExistDelegatedStakingConds(ctx, &appdelegatedstakingmwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appdelegatedstaking")
	}
	return nil
}
