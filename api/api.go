package api

import (
	"context"

	"github.com/NpoolPlatform/good-gateway/api/deviceinfo"

	v1 "github.com/NpoolPlatform/message/npool/good/gw/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterGatewayServer(server, &Server{})
	deviceinfo.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := deviceinfo.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
