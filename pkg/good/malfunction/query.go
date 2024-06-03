package malfunction

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/malfunction"
	malfunctionmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

type queryHandler struct {
	*Handler
	malfunctions           []*malfunctionmwpb.Malfunction
	compensateOrderNumbers map[string]uint32
	infos                  []*npool.Malfunction
}

func (h *queryHandler) getCompensateOrderNumbers(ctx context.Context) (err error) {
	h.compensateOrderNumbers, err = goodgwcommon.GetCompensateOrderNumbers(ctx, func() (compensateFromIDs []string) {
		for _, malfunction := range h.malfunctions {
			compensateFromIDs = append(compensateFromIDs, malfunction.EntID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CompensatedOrders = h.compensateOrderNumbers[info.EntID]
	}
}

func (h *Handler) GetMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	info, err := malfunctionmwcli.GetMalfunction(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid malfunction")
	}

	handler := &queryHandler{
		malfunctions: []*malfunctionmwpb.Malfunction{info},
	}
	if err := handler.getCompensateOrderNumbers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetMalfunctions(ctx context.Context) ([]*npool.Malfunction, uint32, error) {
	conds := &malfunctionmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	infos, total, err := malfunctionmwcli.GetMalfunctions(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		malfunctions: infos,
	}

	if err := handler.getCompensateOrderNumbers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, total, nil
}
