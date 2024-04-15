package poster

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/poster"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	poster.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	poster.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := poster.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
