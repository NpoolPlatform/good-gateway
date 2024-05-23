package api

import (
	"context"

	appfee "github.com/NpoolPlatform/good-gateway/api/app/fee"
	appgood "github.com/NpoolPlatform/good-gateway/api/app/good"
	"github.com/NpoolPlatform/good-gateway/api/app/good/comment"
	default1 "github.com/NpoolPlatform/good-gateway/api/app/good/default"
	"github.com/NpoolPlatform/good-gateway/api/app/good/description"
	displaycolor "github.com/NpoolPlatform/good-gateway/api/app/good/display/color"
	displayname "github.com/NpoolPlatform/good-gateway/api/app/good/display/name"
	"github.com/NpoolPlatform/good-gateway/api/app/good/label"
	"github.com/NpoolPlatform/good-gateway/api/app/good/like"
	"github.com/NpoolPlatform/good-gateway/api/app/good/poster"
	"github.com/NpoolPlatform/good-gateway/api/app/good/recommend"
	appgoodrequired "github.com/NpoolPlatform/good-gateway/api/app/good/required"
	"github.com/NpoolPlatform/good-gateway/api/app/good/score"
	"github.com/NpoolPlatform/good-gateway/api/app/good/topmost"
	topmostconstraint "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/constraint"
	topmostgood "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/good"
	topmostgoodconstraint "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/good/constraint"
	topmostgoodposter "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/good/poster"
	topmostposter "github.com/NpoolPlatform/good-gateway/api/app/good/topmost/poster"
	apppowerrental "github.com/NpoolPlatform/good-gateway/api/app/powerrental"
	apppowerrentalsimulate "github.com/NpoolPlatform/good-gateway/api/app/powerrental/simulate"
	devicetype "github.com/NpoolPlatform/good-gateway/api/device"
	fee "github.com/NpoolPlatform/good-gateway/api/fee"
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
	description.Register(server)
	displayname.Register(server)
	displaycolor.Register(server)
	like.Register(server)
	label.Register(server)
	poster.Register(server)
	recommend.Register(server)
	required.Register(server)
	appgoodrequired.Register(server)
	history.Register(server)
	score.Register(server)
	appgood.Register(server)
	appfee.Register(server)
	fee.Register(server)
	default1.Register(server)
	topmost.Register(server)
	topmostconstraint.Register(server)
	topmostposter.Register(server)
	topmostgood.Register(server)
	topmostgoodconstraint.Register(server)
	topmostgoodposter.Register(server)
	apppowerrentalsimulate.Register(server)
	apppowerrental.Register(server)
}

//nolint:gocyclo,funlen
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
	if err := description.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := displayname.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := displaycolor.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := like.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := label.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := poster.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := recommend.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := required.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgoodrequired.RegisterGateway(mux, endpoint, opts); err != nil {
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
	if err := fee.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appfee.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := default1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmost.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostconstraint.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostposter.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgoodconstraint.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgoodposter.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := apppowerrentalsimulate.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := apppowerrental.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
