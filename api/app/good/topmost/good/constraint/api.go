package constraint

import (
	"context"

	constraint "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	constraint.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	constraint.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := constraint.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
