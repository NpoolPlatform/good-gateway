package fee

import (
	"context"

	fee "github.com/NpoolPlatform/message/npool/good/gw/v1/fee"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fee.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fee.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := fee.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
