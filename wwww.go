package grpcserverfx

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NNewGRPCServer(
	config GRPCServerConfigurer,
	serverMetrics PrometheusServerMetrics,
) *GRPCServer {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			serverMetrics.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			serverMetrics.StreamServerInterceptor(),
		),
	)
	if config.GRPCServerReflection() {
		reflection.Register(srv)
	}
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)
	reg := NewPrometheusRegistry()
	reg.MustRegister(serverMetrics)
	serverMetrics.InitializeMetrics(srv)
	return &GRPCServer{
		srv: srv,
		reg: reg,
	}
}

// API ERROR: This model's maximum context length is 4097 tokens. However, you requested 4171 tokens (171 in the messages, 4000 in the completion). Please reduce the length of the messages or completion.
