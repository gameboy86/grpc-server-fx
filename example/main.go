package main

import (
	grpcserverfx "github.com/gameboy86/grpc-server-fx"
	"github.com/gameboy86/grpc-server-fx/example/services"
	"go.uber.org/fx"
)

type Config struct{}

func (c *Config) GRPCServerPort() int {
	return 8000
}

func (c *Config) GRPCServerReflection() bool {
	return true
}

func NewConfig() *Config {
	return &Config{}
}

func main() {
	fx.New(
		grpcserverfx.Module,
		fx.Provide(
			fx.Annotate(
				NewConfig,
				fx.As(new(grpcserverfx.GRPCServerConfigurer)),
			),
			grpcserverfx.AsService(
				services.NewHelloService,
			),
		),
	).Run()
}
