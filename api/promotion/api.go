package promotion

import (
	"context"

	"github.com/NpoolPlatform/message/npool/good/gw/v1/promotion"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	promotion.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	promotion.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := promotion.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
