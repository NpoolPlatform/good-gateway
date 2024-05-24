package powerrental

import (
	"context"

	powerrental "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	powerrental.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	powerrental.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := powerrental.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
