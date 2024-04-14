package constraint

import (
	"context"

	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/constraint"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkConstraint(ctx); err != nil {
		return nil, err
	}

	if err := topmostconstraintmwcli.DeleteTopMostConstraint(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}

	return h.GetConstraint(ctx)
}
