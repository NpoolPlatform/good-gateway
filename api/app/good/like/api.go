package like

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/like"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	like.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	like.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := like.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
