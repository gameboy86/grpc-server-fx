package grpcserverfx

import (
	"context"

	// "fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GRPCServerConfigurer interface {
	GRPCServerPort() int
	GRPCServerReflection() bool
}

// func NewGRPCHealthServer(srv *grpc.Server)

func NewGRPCServer(config GRPCServerConfigurer) *grpc.Server {
	srv := grpc.NewServer()
	if config.GRPCServerReflection() {
		reflection.Register(srv)
	}
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)
	return srv
}

func New(
	lifecycle fx.Lifecycle,
	listener net.Listener,
	srv GRPCServerer,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.Serve(listener); err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				srv.GracefulStop()
				return nil
			},
		},
	)
	// return srv
}
