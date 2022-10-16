//nolint:nolintlint,dupl
package vendorlocation

import (
	"context"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/vendorlocation"

	"github.com/google/uuid"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/vendorlocation"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"
)

// nolint
func (s *Server) CreateVendorLocation(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocation")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetCountry() == "" {
		logger.Sugar().Errorw("validate", "Country", in.GetCountry())
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Country is empty")
	}

	if in.GetProvince() == "" {
		logger.Sugar().Errorw("validate", "Province", in.GetProvince())
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Province is empty")
	}

	if in.GetCity() == "" {
		logger.Sugar().Errorw("validate", "City", in.GetCity())
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "City is empty")
	}

	if in.GetAddress() == "" {
		logger.Sugar().Errorw("validate", "Address", in.GetAddress())
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Address is empty")
	}

	span = commontracer.TraceInvoker(span, "VendorLocation", "mw", "CreateVendorLocation")

	info, err := mgrcli.CreateVendorLocation(ctx, &mgrpb.VendorLocationReq{
		Country:  &in.Country,
		Province: &in.Province,
		City:     &in.City,
		Address:  &in.Address,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateVendorLocation", "error", err)
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateVendorLocationResponse{
		Info: info,
	}, nil
}

func (s *Server) GetVendorLocation(ctx context.Context, in *npool.GetVendorLocationRequest) (*npool.GetVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocation")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetVendorLocation", "CoinTypeID", in.GetID(), "error", err)
		return &npool.GetVendorLocationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "VendorLocation", "mw", "CreateVendorLocation")

	info, err := mgrcli.GetVendorLocation(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetVendorLocation", "error", err)
		return &npool.GetVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetVendorLocationResponse{
		Info: info,
	}, nil
}

func (s *Server) GetVendorLocations(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateVendorLocation")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "VendorLocation", "mw", "CreateVendorLocation")

	infos, total, err := mgrcli.GetVendorLocations(ctx, nil, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetVendorLocation", "error", err)
		return &npool.GetVendorLocationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetVendorLocationsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateVendorLocation(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateVendorLocation")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateVendorLocation", "ID", in.GetID(), "error", err)
		return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Country != nil {
		if in.GetCountry() == "" {
			logger.Sugar().Errorw("validate", "Country", in.GetCountry())
			return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Country is empty")
		}
	}

	if in.Province != nil {
		if in.GetProvince() == "" {
			logger.Sugar().Errorw("validate", "Province", in.GetProvince())
			return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Province is empty")
		}
	}

	if in.Province != nil {
		if in.GetCity() == "" {
			logger.Sugar().Errorw("validate", "City", in.GetCity())
			return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "City is empty")
		}
	}

	if in.Address != nil {
		if in.GetAddress() == "" {
			logger.Sugar().Errorw("validate", "Address", in.GetAddress())
			return &npool.UpdateVendorLocationResponse{}, status.Error(codes.InvalidArgument, "Address is empty")
		}
	}

	info, err := mgrcli.UpdateVendorLocation(ctx, &mgrpb.VendorLocationReq{
		ID:       &in.ID,
		Country:  in.Country,
		Province: in.Province,
		City:     in.City,
		Address:  in.Address,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateVendorLocation", "error", err)
		return &npool.UpdateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateVendorLocationResponse{
		Info: info,
	}, nil
}
