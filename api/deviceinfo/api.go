package deviceinfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/deviceinfo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	deviceinfo.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	deviceinfo.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := deviceinfo.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
