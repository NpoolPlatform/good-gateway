package api

import (
	"context"

	"github.com/NpoolPlatform/good-gateway/api/appdefaultgood"

	"github.com/NpoolPlatform/good-gateway/api/promotion"
	"github.com/NpoolPlatform/good-gateway/api/recommend"

	"github.com/NpoolPlatform/good-gateway/api/deviceinfo"
	"github.com/NpoolPlatform/good-gateway/api/good"
	"github.com/NpoolPlatform/good-gateway/api/subgood"
	"github.com/NpoolPlatform/good-gateway/api/vendorlocation"
	v1 "github.com/NpoolPlatform/message/npool/good/gw/v1"

	appgood "github.com/NpoolPlatform/good-gateway/api/appgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterGatewayServer(server, &Server{})
	appgood.Register(server)
	deviceinfo.Register(server)
	good.Register(server)
	subgood.Register(server)
	vendorlocation.Register(server)
	promotion.Register(server)
	recommend.Register(server)
	appdefaultgood.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := deviceinfo.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := good.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := subgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := vendorlocation.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := promotion.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := recommend.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appdefaultgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
