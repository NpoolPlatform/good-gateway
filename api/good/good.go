//nolint:nolintlint,dupl
package good

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
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good"

	goodm "github.com/NpoolPlatform/good-gateway/pkg/good"

	"github.com/google/uuid"

	deviceinfocli "github.com/NpoolPlatform/good-manager/pkg/client/deviceinfo"
	vendorlocationcli "github.com/NpoolPlatform/good-manager/pkg/client/vendorlocation"
)

// nolint
func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateGood", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetDurationDays() <= 0 {
		logger.Sugar().Errorw("CreateGood", "DurationDays", in.GetDurationDays())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "DurationDays is invalid")
	}

	if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateGood", "Price", in.GetPrice(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	if in.GetUnit() == "" {
		logger.Sugar().Errorw("CreateGood", "Unit", in.GetUnit(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Unit is empty")
	}

	switch in.GetBenefitType() {
	case mgrpb.BenefitType_BenefitTypePlatform:
	case mgrpb.BenefitType_BenefitTypePool:
	default:
		logger.Sugar().Errorw("CreateGood", "BenefitType", in.GetBenefitType())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "BenefitType is invalid")
	}

	switch in.GetGoodType() {
	case mgrpb.GoodType_GoodTypeClassicMining:
	case mgrpb.GoodType_GoodTypeUnionMining:
	case mgrpb.GoodType_GoodTypeTechniqueFee:
	case mgrpb.GoodType_GoodTypeElectricityFee:
	default:
		logger.Sugar().Errorw("CreateGood", "GoodType", in.GetGoodType())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodType is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("CreateGood", "Title", in.GetTitle())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
	}

	if in.GetUnitAmount() <= 0 {
		logger.Sugar().Errorw("CreateGood", "UnitAmount", in.GetUnitAmount())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "UnitAmount is invalid")
	}

	for _, coinTypeID := range in.GetSupportCoinTypeIDs() {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			logger.Sugar().Errorw("CreateGood", "SupportCoinTypeIDs", in.GetSupportCoinTypeIDs(), "error", err)
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetTotal() <= 0 {
		logger.Sugar().Errorw("CreateGood", "Total", in.GetTotal())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Total is invalid")
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateGood")

	exist, err := deviceinfocli.ExistDeviceInfo(ctx, in.GetDeviceInfoID())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "DeviceInfoID", in.GetDeviceInfoID())
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateGood", "Total", in.GetTotal())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "DeviceInfoID is not exist")
	}

	exist, err = vendorlocationcli.ExistVendorLocation(ctx, in.GetVendorLocationID())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "VendorLocationID", in.GetVendorLocationID())
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("CreateGood", "VendorLocationID", in.GetVendorLocationID())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "VendorLocationID is not exist")
	}

	coinMap, err := goodm.GetCoinType(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "VendorLocationID", in.GetVendorLocationID())
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	_, ok := coinMap[in.GetCoinTypeID()]
	if !ok {
		logger.Sugar().Errorw("CreateGood", "CoinTypeID", in.GetCoinTypeID())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "CoinTypeID is not exist")
	}

	for _, val := range in.GetSupportCoinTypeIDs() {
		_, ok := coinMap[val]
		if !ok {
			logger.Sugar().Errorw("CreateGood", "SupportCoinTypeIDs", in.GetSupportCoinTypeIDs())
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "SupportCoinTypeIDs is not exist")
		}
	}

	info, err := goodm.CreateGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGood(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetGood", "CoinTypeID", in.GetID(), "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateGood")

	info, err := goodm.GetGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetGood", "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateGood")

	infos, total, err := goodm.GetGoods(ctx, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetGood", "error", err)
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

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateGood", "ID", in.GetID(), "error", err)
		return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.DeviceInfoID != nil {
		if _, err := uuid.Parse(in.GetDeviceInfoID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "DeviceInfoID", in.GetDeviceInfoID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		exist, err := deviceinfocli.ExistDeviceInfo(ctx, in.GetDeviceInfoID())
		if err != nil {
			logger.Sugar().Errorw("UpdateGood", "DeviceInfoID", in.GetDeviceInfoID())
			return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
		}

		if !exist {
			logger.Sugar().Errorw("UpdateGood", "Total", in.GetTotal())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DeviceInfoID is not exist")
		}

	}

	if in.DurationDays != nil && in.GetDurationDays() <= 0 {
		logger.Sugar().Errorw("UpdateGood", "DurationDays", in.GetDurationDays())
		return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DurationDays is invalid")
	}

	coinMap, err := goodm.GetCoinType(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateGood", "VendorLocationID", in.GetVendorLocationID())
		return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.CoinTypeID != nil {
		if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "CoinTypeID", in.GetCoinTypeID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		_, ok := coinMap[in.GetCoinTypeID()]
		if !ok {
			logger.Sugar().Errorw("UpdateGood", "CoinTypeID", in.GetCoinTypeID())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "CoinTypeID is not exist")
		}
	}

	if in.InheritFromGoodID != nil {
		if _, err := uuid.Parse(in.GetInheritFromGoodID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "InheritFromGoodID", in.GetInheritFromGoodID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.VendorLocationID != nil {
		if _, err := uuid.Parse(in.GetVendorLocationID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "VendorLocationID", in.GetVendorLocationID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		exist, err := vendorlocationcli.ExistVendorLocation(ctx, in.GetVendorLocationID())
		if err != nil {
			logger.Sugar().Errorw("UpdateGood", "VendorLocationID", in.GetVendorLocationID())
			return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
		}

		if !exist {
			logger.Sugar().Errorw("CreateGood", "VendorLocationID", in.GetVendorLocationID())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "VendorLocationID is not exist")
		}
	}

	if in.Price != nil {
		if price, err := decimal.NewFromString(in.GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetPrice(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	if in.Title != nil {
		if in.GetTitle() == "" {
			logger.Sugar().Errorw("UpdateGood", "Title", in.GetTitle())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
		}
	}

	if in.UnitAmount != nil {
		if in.GetUnitAmount() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "UnitAmount", in.GetUnitAmount())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "UnitAmount is invalid")
		}
	}

	for _, coinTypeID := range in.GetSupportCoinTypeIDs() {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			logger.Sugar().Errorw("UpdateGood", "SupportCoinTypeIDs", in.GetSupportCoinTypeIDs(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		_, ok := coinMap[coinTypeID]
		if !ok {
			logger.Sugar().Errorw("UpdateGood", "SupportCoinTypeIDs", in.GetSupportCoinTypeIDs())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "SupportCoinTypeIDs is not exist")
		}
	}

	if in.Total != nil {
		if in.GetTotal() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Total", in.GetTotal())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Total is invalid")
		}
	}

	info, err := goodm.UpdateGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "error", err)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}
