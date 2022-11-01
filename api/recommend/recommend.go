//nolint:nolintlint,dupl
package recommend

import (
	"context"
	"fmt"

	appmgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"
	appusermgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/appuser"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/recommend"

	appgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/recommend"

	npoolpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"

	recommendm "github.com/NpoolPlatform/good-gateway/pkg/recommend"
)

// nolint
func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetRecommenderID()); err != nil {
		logger.Sugar().Errorw("validate", "RecommenderID", in.GetRecommenderID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("RecommenderID is invalid: %v", err))
	}

	if in.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
	}

	exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
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
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, "App Good is not found")
	}

	exist, err = appusermgrcli.ExistAppUser(ctx, in.GetRecommenderID())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, "Recommender is not found")
	}

	span = commontracer.TraceInvoker(span, "Recommend", "mw", "CreateRecommend")

	info, err := recommendm.CreateRecommend(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateRecommend", "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) CreateAppRecommend(ctx context.Context, in *npool.CreateAppRecommendRequest) (*npool.CreateAppRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetAppID", in.GetTargetAppID(), "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("TargetAppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetRecommenderID()); err != nil {
		logger.Sugar().Errorw("validate", "RecommenderID", in.GetRecommenderID(), "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("RecommenderID is invalid: %v", err))
	}

	if in.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
	}

	app, err := appmgrcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if app == nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, "App is not exist")
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
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	exist, err = appusermgrcli.ExistAppUser(ctx, in.GetRecommenderID())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.InvalidArgument, "Recommender is not found")
	}

	span = commontracer.TraceInvoker(span, "Recommend", "mw", "CreateRecommend")

	info, err := recommendm.CreateRecommend(ctx, &npool.CreateRecommendRequest{
		AppID:          in.TargetAppID,
		GoodID:         in.GoodID,
		RecommenderID:  in.RecommenderID,
		Message:        in.Message,
		RecommendIndex: in.RecommendIndex,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateRecommend", "error", err)
		return &npool.CreateAppRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRecommendResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRecommends(ctx context.Context, in *npool.GetRecommendsRequest) (*npool.GetRecommendsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.GetRecommendsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "Recommend", "mgr", "GetRecommends")

	infos, total, err := recommendm.GetRecommends(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetRecommend", "error", err)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRecommendsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppRecommends(ctx context.Context, in *npool.GetAppRecommendsRequest) (*npool.GetAppRecommendsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRecommends")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.GetAppRecommendsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "Recommend", "mgr", "GetAppRecommends")

	infos, total, err := recommendm.GetRecommends(ctx, in.GetTargetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetRecommend", "error", err)
		return &npool.GetAppRecommendsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRecommendsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("ID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GetAppID is invalid: %v", err))
	}

	recommend, err := mgrcli.GetRecommend(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	if recommend.AppID != in.GetAppID() {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	if in.Message != nil {
		if in.GetMessage() == "" {
			logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
			return &npool.UpdateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
		}
	}

	info, err := recommendm.UpdateRecommend(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("UpdateRecommend", "error", err)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRecommendResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateAppRecommend(ctx context.Context, in *npool.UpdateAppRecommendRequest) (*npool.UpdateAppRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.UpdateAppRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GetAppID is invalid: %v", err))
	}

	app, err := appmgrcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateAppRecommendResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if app == nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.UpdateAppRecommendResponse{}, status.Error(codes.InvalidArgument, "App is not exist")
	}

	if in.Message != nil {
		if in.GetMessage() == "" {
			logger.Sugar().Errorw("validate", "Message", in.GetMessage(), "error", err)
			return &npool.UpdateAppRecommendResponse{}, status.Error(codes.InvalidArgument, "Message is empty")
		}
	}

	info, err := recommendm.UpdateRecommend(ctx, &npool.UpdateRecommendRequest{
		ID:             in.ID,
		AppID:          in.TargetAppID,
		Message:        in.Message,
		RecommendIndex: in.RecommendIndex,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateRecommend", "error", err)
		return &npool.UpdateAppRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRecommendResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteRecommend(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("ID is invalid: %v", err))
	}

	info, err := mgrcli.DeleteRecommend(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("UpdateRecommend", "error", err)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRecommendResponse{
		Info: &npool.Recommend{
			ID:             info.ID,
			AppID:          info.AppID,
			GoodID:         info.GoodID,
			RecommenderID:  info.RecommenderID,
			Message:        info.Message,
			RecommendIndex: info.RecommendIndex,
			CreatedAt:      info.CreatedAt,
			UpdatedAt:      info.UpdatedAt,
		},
	}, nil
}
