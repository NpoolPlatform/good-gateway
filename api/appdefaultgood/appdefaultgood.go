//nolint:nolintlint,dupl
package appdefaultgood

import (
	"context"
	"fmt"

	appgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
	appgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	appdefaultgoodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/appdefaultgood"

	"github.com/google/uuid"

	appdefaultgoodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appdefaultgood"

	appcoininfocli "github.com/NpoolPlatform/chain-middleware/pkg/client/appcoin"
	appcoininfopb "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"
)

func validate(ctx context.Context, in *npool.CreateAppDefaultGoodRequest) error {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	coin, err := appcoininfocli.GetCoinOnly(ctx, &appcoininfopb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		CoinTypeID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetCoinTypeID(),
		},
	})
	if err != nil {
		return err
	}
	if coin == nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "GoodID", in.GetGoodID(), "CoinTypeID", in.GetCoinTypeID())
		return fmt.Errorf("app coin is not exist")
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
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return err
	}
	if !exist {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID())
		return fmt.Errorf("good is not exist")
	}

	exist, err = appdefaultgoodmgrcli.ExistAppDefaultGoodConds(ctx, &appdefaultgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
		CoinTypeID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetCoinTypeID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID(), "error", err)
		return err
	}
	if exist {
		logger.Sugar().Errorw("validate", "GoodID", in.GetGoodID())
		return fmt.Errorf("app default good already exist")
	}
	return nil
}

func (s *Server) CreateNAppDefaultGood(
	ctx context.Context,
	in *npool.CreateNAppDefaultGoodRequest,
) (
	*npool.CreateNAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	err = validate(ctx, &npool.CreateAppDefaultGoodRequest{
		AppID:      in.GetTargetAppID(),
		GoodID:     in.GetGoodID(),
		CoinTypeID: in.GetCoinTypeID(),
	})
	if err != nil {
		logger.Sugar().Errorw("CreateNAppDefaultGood", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.CreateNAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := appdefaultgoodmgrcli.CreateAppDefaultGood(
		ctx,
		&appdefaultgoodmgrpb.AppDefaultGoodReq{
			AppID:      &in.TargetAppID,
			GoodID:     &in.GoodID,
			CoinTypeID: &in.CoinTypeID,
		},
	)
	if err != nil {
		logger.Sugar().Errorw("CreateNAppDefaultGood", "error", err)
		return &npool.CreateNAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNAppDefaultGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppDefaultGood(
	ctx context.Context,
	in *npool.CreateAppDefaultGoodRequest,
) (
	*npool.CreateAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	err = validate(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateNAppDefaultGood", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.CreateAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := appdefaultgoodmgrcli.CreateAppDefaultGood(
		ctx,
		&appdefaultgoodmgrpb.AppDefaultGoodReq{
			AppID:      &in.AppID,
			GoodID:     &in.GoodID,
			CoinTypeID: &in.CoinTypeID,
		},
	)
	if err != nil {
		logger.Sugar().Errorw("CreateAppDefaultGood", "error", err)
		return &npool.CreateAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppDefaultGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppDefaultGoods(
	ctx context.Context,
	in *npool.GetAppDefaultGoodsRequest,
) (
	*npool.GetAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "AppDefaultGood", "mw", "GetAppDefaultGoods")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "AppID", in.GetAppID(), "error", err)
		return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appdefaultgoodmgrcli.GetAppDefaultGoods(ctx, &appdefaultgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
	}, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "error", err)
		return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetNAppDefaultGoods(
	ctx context.Context,
	in *npool.GetNAppDefaultGoodsRequest,
) (
	*npool.GetNAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "AppDefaultGood", "mw", "GetAppDefaultGoods")

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "AppID", in.GetTargetAppID(), "error", err)
		return &npool.GetNAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appdefaultgoodmgrcli.GetAppDefaultGoods(ctx, &appdefaultgoodmgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
	}, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "error", err)
		return &npool.GetNAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNAppDefaultGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) DeleteAppDefaultGood(
	ctx context.Context,
	in *npool.DeleteAppDefaultGoodRequest,
) (
	*npool.DeleteAppDefaultGoodResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	row, err := appdefaultgoodmgrcli.GetAppDefaultGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	if row.GetAppID() != in.GetAppID() {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", "permission denied")
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	info, err := appdefaultgoodmgrcli.DeleteAppDefaultGood(ctx, in.GetID())
	if err != nil {
		return &npool.DeleteAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppDefaultGoodResponse{
		Info: info,
	}, nil
}
