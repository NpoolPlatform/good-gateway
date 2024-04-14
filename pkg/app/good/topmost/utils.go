package topmost

import (
	"context"
	"fmt"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func CheckTopMost(ctx context.Context, appID, topMostID string) error {
	exist, err := topmostmwcli.ExistTopMostConds(ctx, &topmostmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: topMostID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: appID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmost")
	}
	return nil
}
