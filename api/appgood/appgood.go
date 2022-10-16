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

func (s *Server) GetAppGoods(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "AppGood", "mw", "CreateAppGood")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateAppGood", "AppID", in.GetAppID(), "error", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appgood1.GetAppGoods(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAppGood", "error", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodsResponse{
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

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdateAppGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
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
