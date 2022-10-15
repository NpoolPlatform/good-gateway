package appgood

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/appgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appgood.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	appgood.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return appgood.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
