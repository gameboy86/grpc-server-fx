package grpcserverfx

import (
	"net"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GRPCServerConfigurer interface {
	GRPCServerPort() int
	GRPCServerReflection() bool
}

type GRPCServerer interface {
	RegisterService(desc *grpc.ServiceDesc, impl any)
	Serve(lis net.Listener) error
	GracefulStop()
	PrometheusRegistry() *prometheus.Registry
}

type GRPCServer struct {
	srv  *grpc.Server
	prom *PrometheusGRPC
}

func (s *GRPCServer) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.srv.RegisterService(desc, impl)
}

func (s *GRPCServer) Serve(lis net.Listener) error {
	return s.srv.Serve(lis)
}

func (s *GRPCServer) ServeMetrics(lis net.Listener) error {
	return s.srv.Serve(lis)
}

func (s *GRPCServer) GracefulStop() {
	s.srv.GracefulStop()
}

func (s *GRPCServer) PrometheusRegistry() *prometheus.Registry {
	return s.prom.Registry
}

func NewGRPCServer(
	config GRPCServerConfigurer,
	promethesu *PrometheusGRPC,
) *GRPCServer {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			promethesu.ServerMetrics.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			promethesu.ServerMetrics.StreamServerInterceptor(),
		),
	)
	if config.GRPCServerReflection() {
		reflection.Register(srv)
	}
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)

	promethesu.InitializeMetrics(srv)

	return &GRPCServer{
		srv:  srv,
		prom: promethesu,
	}
}
