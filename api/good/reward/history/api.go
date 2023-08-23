package history

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/good/reward/history"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := history.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
