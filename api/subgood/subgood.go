//nolint:nolintlint,dupl
package subgood

import (
	"context"
	"fmt"

	appgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	submgrcli "github.com/NpoolPlatform/good-manager/pkg/client/subgood"
	submgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/subgood"

	"github.com/google/uuid"

	subgoodm "github.com/NpoolPlatform/good-gateway/pkg/subgood"

	npoolpb "github.com/NpoolPlatform/message/npool"
)

// nolint
func (s *Server) CreateSubGood(ctx context.Context, in *npool.CreateSubGoodRequest) (*npool.CreateSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetMainGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("MainGoodID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetSubGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is invalid: %v", err))
	}

	if in.GetMainGoodID() == in.GetSubGoodID() {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is MainGoodID"))
	}

	exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetMainGoodID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	exist, err = appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetSubGoodID(),
		},
		GoodIDs: nil,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	exist, err = submgrcli.ExistSubGoodConds(ctx, &submgrpb.Conds{
		MainGoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetMainGoodID(),
		},
		SubGoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetSubGoodID(),
		},
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if exist {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.InvalidArgument, "Sub Good already exist")
	}

	span = commontracer.TraceInvoker(span, "SubGood", "mw", "CreateSubGood")

	info, err := subgoodm.CreateSubGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateSubGood", "error", err)
		return &npool.CreateSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubGoodResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) CreateAppSubGood(ctx context.Context, in *npool.CreateAppSubGoodRequest) (*npool.CreateAppSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetMainGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("MainGoodID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetSubGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is invalid: %v", err))
	}

	exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetMainGoodID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "MainGoodID", in.GetMainGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	exist, err = appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetSubGoodID(),
		},
		GoodIDs: nil,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
	}

	exist, err = submgrcli.ExistSubGoodConds(ctx, &submgrpb.Conds{
		MainGoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetMainGoodID(),
		},
		SubGoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetSubGoodID(),
		},
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if exist {
		logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, "Sub Good already exist")
	}

	span = commontracer.TraceInvoker(span, "SubGood", "mw", "CreateSubGood")

	info, err := subgoodm.CreateSubGood(ctx, &npool.CreateSubGoodRequest{
		AppID:      in.TargetAppID,
		MainGoodID: in.MainGoodID,
		SubGoodID:  in.SubGoodID,
		Must:       in.Must,
		Commission: in.Commission,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateSubGood", "error", err)
		return &npool.CreateAppSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppSubGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubGoods(ctx context.Context, in *npool.GetSubGoodsRequest) (*npool.GetSubGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.GetSubGoodsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "SubGood", "mgr", "GetSubGoods")

	infos, total, err := subgoodm.GetSubGoods(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetSubGood", "error", err)
		return &npool.GetSubGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppSubGoods(ctx context.Context, in *npool.GetAppSubGoodsRequest) (*npool.GetAppSubGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.GetAppSubGoodsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	span = commontracer.TraceInvoker(span, "SubGood", "mgr", "GetSubGoods")

	infos, total, err := subgoodm.GetSubGoods(ctx, in.GetTargetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetSubGood", "error", err)
		return &npool.GetAppSubGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppSubGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateSubGood(ctx context.Context, in *npool.UpdateSubGoodRequest) (*npool.UpdateSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSubGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSubGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	subGood, err := submgrcli.GetSubGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("UpdateSubGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	if subGood.GetAppID() != in.AppID {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", "subgood appid != appid")
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid"))
	}

	if in.SubGoodID != nil {
		if _, err := uuid.Parse(in.GetSubGoodID()); err != nil {
			logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
			return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is invalid: %v", err))
		}

		exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetAppID(),
			},
			GoodID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetSubGoodID(),
			},
			GoodIDs: nil,
		})

		if err != nil {
			logger.Sugar().Errorw("validate", "SubGood", in.GetSubGoodID(), "error", err)
			return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		if !exist {
			logger.Sugar().Errorw("validate", "SubGood", in.GetSubGoodID(), "error", err)
			return &npool.UpdateSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
		}
	}

	info, err := subgoodm.UpdateSubGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("UpdateSubGood", "error", err)
		return &npool.UpdateSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSubGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateAppSubGood(ctx context.Context, in *npool.UpdateAppSubGoodRequest) (*npool.UpdateAppSubGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSubGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSubGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if in.SubGoodID != nil {
		if _, err := uuid.Parse(in.GetSubGoodID()); err != nil {
			logger.Sugar().Errorw("validate", "SubGoodID", in.GetSubGoodID(), "error", err)
			return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("SubGoodID is invalid: %v", err))
		}

		exist, err := appgoodmgrcli.ExistAppGoodConds(ctx, &appgoodmgrpb.Conds{
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetTargetAppID(),
			},
			GoodID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetSubGoodID(),
			},
			GoodIDs: nil,
		})

		if err != nil {
			logger.Sugar().Errorw("validate", "SubGood", in.GetSubGoodID(), "error", err)
			return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		if !exist {
			logger.Sugar().Errorw("validate", "SubGood", in.GetSubGoodID(), "error", err)
			return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.InvalidArgument, "App Good not exist")
		}
	}

	info, err := subgoodm.UpdateSubGood(ctx, &npool.UpdateSubGoodRequest{
		ID:         in.ID,
		SubGoodID:  in.SubGoodID,
		Must:       in.Must,
		Commission: in.Commission,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateSubGood", "error", err)
		return &npool.UpdateAppSubGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppSubGoodResponse{
		Info: info,
	}, nil
}
