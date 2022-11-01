//nolint:nolintlint,dupl
package promotion

import (
	"context"
	"fmt"
	"time"

	appgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/promotion"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/promotion"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/promotion"

	npoolpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"

	promotionm "github.com/NpoolPlatform/good-gateway/pkg/promotion"

	appmgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"
)

// nolint
func (s *Server) CreatePromotion(ctx context.Context, in *npool.CreatePromotionRequest) (*npool.CreatePromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotion")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if in.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
	}

	now := uint32(time.Now().Unix())
	if in.GetStartAt() <= now {
		logger.Sugar().Errorw("validate", "StartAt", in.GetStartAt(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("StartAt is invalid"))
	}

	if in.GetEndAt() <= in.GetStartAt() {
		logger.Sugar().Errorw("validate", "EndAt", in.GetEndAt(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("EndAt is invalid"))
	}

	if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateGood", "Price", in.GetPrice(), "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	exist, err := mgrcli.ExistPromotionConds(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	if exist {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.InvalidArgument, "Promotion already exists")
	}

	span = commontracer.TraceInvoker(span, "Promotion", "mw", "CreatePromotion")

	info, err := promotionm.CreatePromotion(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreatePromotion", "error", err)
		return &npool.CreatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePromotionResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) CreateAppPromotion(ctx context.Context, in *npool.CreateAppPromotionRequest) (*npool.CreateAppPromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePromotion")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetAppID", in.GetTargetAppID(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("TargetAppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if in.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
	}

	now := uint32(time.Now().Unix())
	if in.GetStartAt() <= now {
		logger.Sugar().Errorw("validate", "StartAt", in.GetStartAt(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("StartAt is invalid"))
	}

	if in.GetEndAt() <= in.GetStartAt() {
		logger.Sugar().Errorw("validate", "EndAt", in.GetEndAt(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("EndAt is invalid"))
	}

	if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateGood", "Price", in.GetPrice(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	exist, err := mgrcli.ExistPromotionConds(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
		EndAt: &npoolpb.Uint32Val{
			Op:    cruder.GT,
			Value: in.GetStartAt(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	if exist {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "Promotion already exists")
	}

	exist, err = appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
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
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	app, err := appmgrcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if app == nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "App is not exist")
	}

	span = commontracer.TraceInvoker(span, "Promotion", "mw", "CreatePromotion")

	info, err := promotionm.CreatePromotion(ctx, &npool.CreatePromotionRequest{
		AppID:   in.TargetAppID,
		GoodID:  in.GoodID,
		Message: in.Message,
		StartAt: in.StartAt,
		EndAt:   in.EndAt,
		Price:   in.Price,
		Posters: in.Posters,
	})
	if err != nil {
		logger.Sugar().Errorw("CreatePromotion", "error", err)
		return &npool.CreateAppPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppPromotionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPromotions(ctx context.Context, in *npool.GetPromotionsRequest) (*npool.GetPromotionsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.GetPromotionsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "Promotion", "mgr", "GetPromotions")

	infos, total, err := promotionm.GetPromotions(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetPromotion", "error", err)
		return &npool.GetPromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPromotionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppPromotions(ctx context.Context, in *npool.GetAppPromotionsRequest) (*npool.GetAppPromotionsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPromotions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetAppID", in.GetTargetAppID(), "error", err)
		return &npool.GetAppPromotionsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "Promotion", "mgr", "GetPromotions")

	infos, total, err := promotionm.GetPromotions(ctx, in.GetTargetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetPromotion", "error", err)
		return &npool.GetAppPromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppPromotionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdatePromotion(ctx context.Context, in *npool.UpdatePromotionRequest) (*npool.UpdatePromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdatePromotion")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("ID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GetAppID is invalid: %v", err))
	}

	promotion, err := mgrcli.GetPromotion(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	if promotion.AppID != in.GetAppID() {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	if in.Message != nil {
		if in.GetMessage() == "" {
			logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
			return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
		}
	}

	if in.StartAt != nil {
		now := uint32(time.Now().Unix())
		if in.GetStartAt() <= now {
			logger.Sugar().Errorw("validate", "StartAt", in.GetStartAt(), "error", err)
			return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("StartAt is invalid"))
		}

		if in.GetEndAt() <= in.GetStartAt() {
			logger.Sugar().Errorw("validate", "EndAt", in.GetEndAt(), "error", err)
			return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("EndAt is invalid"))
		}
	}

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("CreateGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdatePromotionResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	info, err := promotionm.UpdatePromotion(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("UpdatePromotion", "error", err)
		return &npool.UpdatePromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePromotionResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateAppPromotion(ctx context.Context, in *npool.UpdateAppPromotionRequest) (*npool.UpdateAppPromotionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdatePromotion")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if in.Message != nil {
		if in.GetMessage() == "" {
			logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
			return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "Message is empty")
		}
	}

	if in.StartAt != nil {
		now := uint32(time.Now().Unix())
		if in.GetStartAt() <= now {
			logger.Sugar().Errorw("validate", "StartAt", in.GetStartAt(), "error", err)
			return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "StartAt is invalid")
		}

		if in.GetEndAt() <= in.GetStartAt() {
			logger.Sugar().Errorw("validate", "EndAt", in.GetEndAt(), "error", err)
			return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "EndAt is invalid")
		}
	}

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("CreateGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	app, err := appmgrcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if app == nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateAppPromotionResponse{}, status.Error(codes.InvalidArgument, "App is not exist")
	}

	info, err := promotionm.UpdatePromotion(ctx, &npool.UpdatePromotionRequest{
		ID:      in.ID,
		AppID:   in.TargetAppID,
		Message: in.Message,
		StartAt: in.StartAt,
		EndAt:   in.EndAt,
		Price:   in.Price,
		Posters: in.Posters,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdatePromotion", "error", err)
		return &npool.UpdateAppPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppPromotionResponse{
		Info: info,
	}, nil
}
