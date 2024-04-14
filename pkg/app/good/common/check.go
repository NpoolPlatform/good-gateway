package common

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

type CheckHandler struct {
	goodgwcommon.AppUserCheckHandler
	GoodID    *string
	AppGoodID *string
}

func (h *CheckHandler) CheckGood(ctx context.Context) error {
	exist, err := goodmwcli.ExistGoodConds(ctx, &goodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid good")
	}
	return nil
}

func (h *CheckHandler) CheckAppGoodWithAppGoodID(ctx context.Context, appGoodID string) error {
	exist, err := appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: appGoodID},
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

func (h *CheckHandler) CheckAppGood(ctx context.Context) error {
	return h.CheckAppGoodWithAppGoodID(ctx, *h.AppGoodID)
}
