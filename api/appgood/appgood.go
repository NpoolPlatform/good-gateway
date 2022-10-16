//nolint:nolintlint,dupl
package appgood

import (
	"context"

	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/good"

	appgood1 "github.com/NpoolPlatform/good-gateway/pkg/appgood"

	"github.com/google/uuid"
)

// nolint
func (s *Server) CreateAppGood(ctx context.Context, in *npool.CreateAppGoodRequest) (*npool.CreateAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateAppGood", "AppID", in.GetAppID(), "error", err)
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	price, err := decimal.NewFromString(in.GetPrice())
	if err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateAppGood", "Price", in.GetPrice(), "error", err)
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	if in.GetGoodName() == "" {
		logger.Sugar().Errorw("CreateAppGood", "GoodName", in.GetGoodName())
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
	}

	if in.GetDisplayIndex() < 0 {
		logger.Sugar().Errorw("CreateAppGood", "DisplayIndex", in.GetDisplayIndex())
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
	}

	if in.GetPurchaseLimit() <= 0 {
		logger.Sugar().Errorw("CreateAppGood", "PurchaseLimit", in.GetPurchaseLimit())
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
	}

	if in.GetCommissionPercent() >= 100 || in.GetCommissionPercent() < 0 {
		logger.Sugar().Errorw("CreateAppGood", "CommissionPercent", in.GetCommissionPercent())
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
	}

	exist, err := goodmgrcli.ExistGood(ctx, in.GetGoodID())
	if err != nil {
		logger.Sugar().Errorw("CreateAppGood", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateAppGood", "GoodID", in.GetGoodID(), "exist", exist)
		return &npool.CreateAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodID not exist")
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateAppGood")

	info, err := appgood1.CreateAppGood(
		ctx,
		in.GetAppID(),
		in.GetGoodID(),
		in.GetOnline(),
		in.GetVisible(),
		in.GetGoodName(),
		in.GetPrice(),
		in.GetDisplayIndex(),
		in.GetPurchaseLimit(),
		in.GetCommissionPercent(),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateAppGood", "error", err)
		return &npool.CreateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppGoodResponse{
		Info: info,
	}, nil
}

/*
func (s *Server) GetAppGood(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetGood", "ID", in.GetID(), "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "GetGood")

	info, err := goodmw.GetGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetGood", "ID", in.GetID(), "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetConds().ID != nil {
		if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "ID", in.GetConds().GetID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		exist, err := goodmgrcli.ExistGood(ctx, in.GetConds().GetGoodID().GetValue())
		if err != nil {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if !exist {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "exist", exist)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, "GoodID not exist")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "GetGoods")

	limit := in.GetLimit()
	if limit <= 0 {
		limit = constant1.DefaultLimit
	}

	infos, total, err := goodmw.GetGoods(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "Conds", in.GetConds(), "error", err)
		return &npool.GetGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateGood", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().AppID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "AppID", in.GetInfo().GetAppID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().Price != nil {
		if price, err := decimal.NewFromString(in.GetInfo().GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetInfo().GetPrice(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	if in.GetInfo().GoodName != nil {
		if in.GetInfo().GetGoodName() == "" {
			logger.Sugar().Errorw("UpdateGood", "GoodName", in.GetInfo().GetGoodName())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
		}
	}

	if in.GetInfo().DisplayIndex != nil {
		if in.GetInfo().GetDisplayIndex() < 0 {
			logger.Sugar().Errorw("UpdateGood", "DisplayIndex", in.GetInfo().GetDisplayIndex())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
		}
	}

	if in.GetInfo().PurchaseLimit != nil {
		if in.GetInfo().GetPurchaseLimit() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "PurchaseLimit", in.GetInfo().GetPurchaseLimit())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
		}
	}

	if in.GetInfo().CommissionPercent != nil {
		if in.GetInfo().GetCommissionPercent() >= 100 || in.GetInfo().GetCommissionPercent() < 0 {
			logger.Sugar().Errorw("UpdateGood", "CommissionPercent", in.GetInfo().GetCommissionPercent())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "UpdateGood")

	info, err := goodmw.UpdateGood(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateGood", "Good", in.GetInfo())
		return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}
*/
