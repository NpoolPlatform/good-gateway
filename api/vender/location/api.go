package location

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/vender/location"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	location.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	location.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := location.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
