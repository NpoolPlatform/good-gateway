package subgood

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/subgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	subgood.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	subgood.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := subgood.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
