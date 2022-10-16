//nolint:nolintlint,dupl
package deviceinfo

import (
	"context"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/deviceinfo"

	"github.com/google/uuid"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/deviceinfo"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"
)

// nolint
func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfo")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetType() == "" {
		logger.Sugar().Errorw("validate", "Type", in.GetType())
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "Type is empty")
	}

	if in.GetManufacturer() == "" {
		logger.Sugar().Errorw("validate", "Manufacturer", in.GetManufacturer())
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "Manufacturer is empty")
	}

	if in.GetShipmentAt() < 0 {
		logger.Sugar().Errorw("validate", "ShipmentAt", in.ShipmentAt)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "ShipmentAt is empty")
	}

	span = commontracer.TraceInvoker(span, "DeviceInfo", "mw", "CreateDeviceInfo")

	info, err := mgrcli.CreateDeviceInfo(ctx, &mgrpb.DeviceInfoReq{
		Type:            &in.Type,
		Manufacturer:    &in.Manufacturer,
		PowerComsuption: &in.PowerComsuption,
		ShipmentAt:      &in.ShipmentAt,
		Posters:         in.Posters,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateDeviceInfo", "error", err)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDeviceInfoResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDeviceInfo(ctx context.Context, in *npool.GetDeviceInfoRequest) (*npool.GetDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfo")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetDeviceInfo", "CoinTypeID", in.GetID(), "error", err)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "DeviceInfo", "mw", "CreateDeviceInfo")

	info, err := mgrcli.GetDeviceInfo(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetDeviceInfo", "error", err)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeviceInfoResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDeviceInfos(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDeviceInfo")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "DeviceInfo", "mw", "CreateDeviceInfo")

	infos, total, err := mgrcli.GetDeviceInfos(ctx, nil, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetDeviceInfo", "error", err)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeviceInfosResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateDeviceInfo")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateDeviceInfo", "ID", in.GetID(), "error", err)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Type != nil {
		if in.GetType() == "" {
			logger.Sugar().Errorw("validate", "Type", in.GetType())
			return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "Type is empty")
		}
	}

	if in.Manufacturer != nil {
		if in.GetManufacturer() == "" {
			logger.Sugar().Errorw("validate", "Manufacturer", in.GetManufacturer())
			return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "Manufacturer is empty")
		}
	}

	if in.ShipmentAt != nil {
		if in.GetShipmentAt() < 0 {
			logger.Sugar().Errorw("validate", "ShipmentAt", in.ShipmentAt)
			return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.InvalidArgument, "ShipmentAt is empty")
		}
	}

	info, err := mgrcli.UpdateDeviceInfo(ctx, &mgrpb.DeviceInfoReq{
		Type:            in.Type,
		Manufacturer:    in.Manufacturer,
		PowerComsuption: in.PowerComsuption,
		ShipmentAt:      in.ShipmentAt,
		Posters:         in.Posters,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateDeviceInfo", "error", err)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDeviceInfoResponse{
		Info: info,
	}, nil
}
