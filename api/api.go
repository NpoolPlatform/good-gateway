package api

import (
	"context"

	appgood "github.com/NpoolPlatform/good-gateway/api/app/good"
	default1 "github.com/NpoolPlatform/good-gateway/api/app/good/default"
	"github.com/NpoolPlatform/good-gateway/api/app/good/topmost"
	topmostgood "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/good"
	"github.com/NpoolPlatform/good-gateway/api/deviceinfo"
	"github.com/NpoolPlatform/good-gateway/api/good"
	"github.com/NpoolPlatform/good-gateway/api/good/comment"
	"github.com/NpoolPlatform/good-gateway/api/good/like"
	"github.com/NpoolPlatform/good-gateway/api/good/recommend"
	"github.com/NpoolPlatform/good-gateway/api/good/required"
	"github.com/NpoolPlatform/good-gateway/api/good/reward/history"
	"github.com/NpoolPlatform/good-gateway/api/good/score"
	"github.com/NpoolPlatform/good-gateway/api/vender/brand"
	"github.com/NpoolPlatform/good-gateway/api/vender/location"

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
	brand.Register(server)
	location.Register(server)
	good.Register(server)
	like.Register(server)
	recommend.Register(server)
	required.Register(server)
	history.Register(server)
	score.Register(server)
	appgood.Register(server)
	default1.Register(server)
	topmost.Register(server)
	topmostgood.Register(server)
}

//nolint:gocyclo
func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := deviceinfo.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := brand.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := location.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := good.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := comment.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := like.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := recommend.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := required.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := history.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := score.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := default1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmost.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
