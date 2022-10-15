package api

import (
	"context"

	good "github.com/NpoolPlatform/message/npool/good/gw/v1"

	appgood "github.com/NpoolPlatform/good-gateway/api/appgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	good.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	good.RegisterGatewayServer(server, &Server{})
	appgood.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := good.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
