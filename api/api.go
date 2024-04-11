package api

import (
	"context"

	appgood "github.com/NpoolPlatform/good-gateway/api/app/good"
	"github.com/NpoolPlatform/good-gateway/api/app/good/comment"
	default1 "github.com/NpoolPlatform/good-gateway/api/app/good/default"
	"github.com/NpoolPlatform/good-gateway/api/app/good/like"
	"github.com/NpoolPlatform/good-gateway/api/app/good/recommend"
	"github.com/NpoolPlatform/good-gateway/api/app/good/score"
	"github.com/NpoolPlatform/good-gateway/api/app/good/topmost"
	topmostgood "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/good"
	apppowerrentalsimulate "github.com/NpoolPlatform/good-gateway/api/app/powerrental/simulate"
	devicetype "github.com/NpoolPlatform/good-gateway/api/device"
	"github.com/NpoolPlatform/good-gateway/api/good"
	"github.com/NpoolPlatform/good-gateway/api/good/required"
	"github.com/NpoolPlatform/good-gateway/api/good/reward/history"
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
	devicetype.Register(server)
	brand.Register(server)
	location.Register(server)
	good.Register(server)
	comment.Register(server)
	like.Register(server)
	recommend.Register(server)
	required.Register(server)
	history.Register(server)
	score.Register(server)
	appgood.Register(server)
	default1.Register(server)
	topmost.Register(server)
	topmostgood.Register(server)
	apppowerrentalsimulate.Register(server)
}

//nolint:gocyclo
func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := devicetype.RegisterGateway(mux, endpoint, opts); err != nil {
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
	if err := apppowerrentalsimulate.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
