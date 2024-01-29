package grpcserverfx

import (
	"fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GRPCServerer interface {
	RegisterService(desc *grpc.ServiceDesc, impl any)
	Serve(lis net.Listener) error
	GracefulStop()
}

var Module = fx.Module(
	"grpc_server",
	fx.Provide(
		fx.Annotate(
			NewGRPCServer,
			fx.As(new(GRPCServerer)),
		),
		NewListener,
	),
	fx.Invoke(
		fx.Annotate(
			RegisterServices,
			fx.ParamTags(`group:"service"`),
		),
	),
	fx.Invoke(
		New,
	),
)

var Listener = net.Listen

func NewListener(config GRPCServerConfigurer) (net.Listener, error) {
	ln, err := Listener(
		"tcp",
		fmt.Sprintf(":%v", config.GRPCServerPort()),
	)
	if err != nil {
		return nil, err
	}
	return ln, err
}

func RegisterServices(
	services []Service,
	sr GRPCServerer,
) {
	for _, s := range services {
		sr.RegisterService(s.Description(), s)
	}
}

type Service interface {
	Description() *grpc.ServiceDesc
}

func AsService(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Service)),
		fx.ResultTags(`group:"service"`),
	)
}
