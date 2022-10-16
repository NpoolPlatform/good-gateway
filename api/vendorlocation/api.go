package vendorlocation

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/vendorlocation"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	vendorlocation.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	vendorlocation.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := vendorlocation.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
