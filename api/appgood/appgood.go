//nolint:nolintlint,dupl
package appgood

import (
	"context"

	appmgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

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

	appgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
)

// nolint
func (s *Server) CreateNAppGood(ctx context.Context, in *npool.CreateNAppGoodRequest) (*npool.CreateNAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNAppGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("CreateNAppGood", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	price, err := decimal.NewFromString(in.GetPrice())
	if err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateNAppGood", "Price", in.GetPrice(), "error", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	if in.GetGoodName() == "" {
		logger.Sugar().Errorw("CreateNAppGood", "GoodName", in.GetGoodName())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
	}

	if in.GetDisplayIndex() < 0 {
		logger.Sugar().Errorw("CreateNAppGood", "DisplayIndex", in.GetDisplayIndex())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
	}

	if in.GetPurchaseLimit() <= 0 {
		logger.Sugar().Errorw("CreateNAppGood", "PurchaseLimit", in.GetPurchaseLimit())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
	}

	if in.GetCommissionPercent() >= 100 || in.GetCommissionPercent() < 0 {
		logger.Sugar().Errorw("CreateNAppGood", "CommissionPercent", in.GetCommissionPercent())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
	}

	good, err := goodmgrcli.GetGood(ctx, in.GetGoodID())
	if err != nil {
		logger.Sugar().Errorw("CreateNAppGood", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if good == nil {
		logger.Sugar().Errorw("CreateNAppGood", "GoodID", in.GetGoodID())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodID not exist")
	}

	if in.GetPrice() < good.GetPrice() {
		logger.Sugar().Errorw("CreateNAppGood", "GoodID", in.GetGoodID())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "price greater than platform price")
	}

	exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateNAppGood", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("CreateNAppGood", "GoodID", in.GetGoodID())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "Good is already exist")
	}

	exist, err = appmgrcli.ExistApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("CreateNAppGood", "err", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateNAppGood", "AppID", in.GetTargetAppID())
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "AppID is not exist")
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateNAppGood")

	info, err := appgood1.CreateAppGood(
		ctx,
		in.GetTargetAppID(),
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
		logger.Sugar().Errorw("CreateNAppGood", "error", err)
		return &npool.CreateNAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNAppGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppGoods(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "AppGood", "mw", "GetAppGoods")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetAppGoods", "AppID", in.GetAppID(), "error", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appgood1.GetAppGoods(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAppGoods", "error", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppGood(ctx context.Context, in *npool.GetAppGoodRequest) (*npool.GetAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetAppGood", "AppID", in.GetAppID(), "error", err)
		return &npool.GetAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("GetAppGood", "GoodID", in.GetGoodID(), "error", err)
		return &npool.GetAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "AppGood", "pgk", "GetAppGood")

	info, err := appgood1.GetAppGood(ctx, in.GetAppID(), in.GetGoodID())
	if err != nil {
		logger.Sugar().Errorw("GetAppGood", "error", err)
		return &npool.GetAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNAppGoods(ctx context.Context, in *npool.GetNAppGoodsRequest) (*npool.GetNAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "AppGood", "mw", "GetAppGoods")

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("GetAppGoods", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.GetNAppGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appgood1.GetAppGoods(ctx, in.GetTargetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAppGoods", "error", err)
		return &npool.GetNAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNAppGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateAppGood(ctx context.Context, in *npool.UpdateAppGoodRequest) (*npool.UpdateAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("UpdateGood", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	appGood, err := appgoodmgrcli.GetAppGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("UpdateGood", "error", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if appGood.AppID != in.GetAppID() {
		logger.Sugar().Errorw("UpdateGood", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}

		good, err := goodmgrcli.GetGood(ctx, appGood.GetGoodID())
		if err != nil {
			logger.Sugar().Errorw("UpdateGood", "error", err)
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
		}

		if appGood.AppID != in.GetAppID() {
			logger.Sugar().Errorw("UpdateGood", "AppID", in.GetAppID(), "error", err)
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
		}

		if in.GetPrice() < good.GetPrice() {
			logger.Sugar().Errorw("CreateNAppGood", "Price", in.GetPrice())
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "price greater than platform price")
		}
	}

	if in.GoodName != nil {
		if in.GetGoodName() == "" {
			logger.Sugar().Errorw("UpdateGood", "GoodName", in.GetGoodName())
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
		}
	}

	if in.DisplayIndex != nil {
		if in.GetDisplayIndex() < 0 {
			logger.Sugar().Errorw("UpdateGood", "DisplayIndex", in.GetDisplayIndex())
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
		}
	}

	if in.PurchaseLimit != nil {
		if in.GetPurchaseLimit() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "PurchaseLimit", in.GetPurchaseLimit())
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
		}
	}

	if in.CommissionPercent != nil {
		if in.GetCommissionPercent() >= 100 || in.GetCommissionPercent() < 0 {
			logger.Sugar().Errorw("UpdateGood", "CommissionPercent", in.GetCommissionPercent())
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "UpdateGood")

	info, err := appgood1.UpdateAppGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("GetAppGood", "error", err)
		return &npool.UpdateAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppGoodResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) UpdateNAppGood(ctx context.Context, in *npool.UpdateNAppGoodRequest) (*npool.UpdateNAppGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNAppGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateNAppGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("UpdateNAppGood", "TargetAppID", in.GetTargetAppID(), "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	app, err := appmgrcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if app == nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "App is not exist")
	}

	appGood, err := appgoodmgrcli.GetAppGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("UpdateGood", "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateNAppGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}

		good, err := goodmgrcli.GetGood(ctx, appGood.GetGoodID())
		if err != nil {
			logger.Sugar().Errorw("UpdateNAppGood", "error", err)
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.Internal, err.Error())
		}

		if in.GetPrice() < good.GetPrice() {
			logger.Sugar().Errorw("UpdateNAppGood", "Price", in.GetPrice())
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "price greater than platform price")
		}
	}

	if in.GoodName != nil {
		if in.GetGoodName() == "" {
			logger.Sugar().Errorw("UpdateNAppGood", "GoodName", in.GetGoodName())
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
		}
	}

	if in.DisplayIndex != nil {
		if in.GetDisplayIndex() < 0 {
			logger.Sugar().Errorw("UpdateNAppGood", "DisplayIndex", in.GetDisplayIndex())
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
		}
	}

	if in.PurchaseLimit != nil {
		if in.GetPurchaseLimit() <= 0 {
			logger.Sugar().Errorw("UpdateNAppGood", "PurchaseLimit", in.GetPurchaseLimit())
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
		}
	}

	if in.CommissionPercent != nil {
		if in.GetCommissionPercent() >= 100 || in.GetCommissionPercent() < 0 {
			logger.Sugar().Errorw("UpdateNAppGood", "CommissionPercent", in.GetCommissionPercent())
			return &npool.UpdateNAppGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "UpdateGood")

	info, err := appgood1.UpdateAppGood(ctx, &npool.UpdateAppGoodRequest{
		ID:                in.ID,
		Online:            in.Online,
		Visible:           in.Visible,
		GoodName:          in.GoodName,
		Price:             in.Price,
		DisplayIndex:      in.DisplayIndex,
		PurchaseLimit:     in.PurchaseLimit,
		CommissionPercent: in.CommissionPercent,
	})
	if err != nil {
		logger.Sugar().Errorw("GetAppGood", "error", err)
		return &npool.UpdateNAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNAppGoodResponse{
		Info: info,
	}, nil
}
