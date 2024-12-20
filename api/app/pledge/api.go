package pledge

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/app/pledge"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	pledge.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	pledge.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := pledge.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
