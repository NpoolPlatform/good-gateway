package common

import (
	"context"
	"fmt"

	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
)

type CheckHandler struct {
	AppID     *string
	AppGoodID *string
}

func (h *CheckHandler) CheckAppGood(ctx context.Context) error {
	exist, err := appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appgood")
	}
	return nil
}
