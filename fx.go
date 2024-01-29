package grpcserverfx

import (
	"go.uber.org/fx"
)

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
			RegisterGRPCServices,
			fx.ParamTags(`group:"service"`),
		),
	),
	fx.Invoke(
		RunServer,
	),
)

func AsService(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(GRPCService)),
		fx.ResultTags(`group:"service"`),
	)
}
