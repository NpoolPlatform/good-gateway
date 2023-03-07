package appdefaultgood

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/appdefaultgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appdefaultgood.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	appdefaultgood.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return appdefaultgood.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
