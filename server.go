package grpcserverfx

import (
	"context"
	"fmt"
	"net"
	"net/http"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

type GRPCServerer interface {
	RegisterService(desc *grpc.ServiceDesc, impl any)
	Serve(lis net.Listener) error
	GracefulStop()
	PrometheusRegistry() *prometheus.Registry
}

type GRPCServer struct {
	srv *grpc.Server
	reg *prometheus.Registry
}

func (s *GRPCServer) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.srv.RegisterService(desc, impl)
}

func (s *GRPCServer) Serve(lis net.Listener) error {
	return s.srv.Serve(lis)
}

func (s *GRPCServer) GracefulStop() {
	s.srv.GracefulStop()
}

func (s *GRPCServer) PrometheusRegistry() *prometheus.Registry {
	return s.reg
}

type GRPCService interface {
	Description() *grpc.ServiceDesc
}

func RegisterGRPCServices(
	services []GRPCService,
	sr GRPCServerer,
) {
	for _, s := range services {
		sr.RegisterService(s.Description(), s)
	}
}

func NewPrometheusRegistry() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func NewGRPCServer(
	config GRPCServerConfigurer,
) *GRPCServer {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets(
				[]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120},
			),
		),
	)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			srvMetrics.StreamServerInterceptor(),
		),
	)
	if config.GRPCServerReflection() {
		reflection.Register(srv)
	}
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)
	reg := NewPrometheusRegistry()
	reg.MustRegister(srvMetrics)
	srvMetrics.InitializeMetrics(srv)
	return &GRPCServer{
		srv: srv,
		reg: reg,
	}
}

func NewListener(config GRPCServerConfigurer) (net.Listener, error) {
	ln, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%v", config.GRPCServerPort()),
	)
	if err != nil {
		return nil, err
	}
	return ln, err
}

func RunServer(
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

				go func() error {
					httpSrv := &http.Server{Addr: ":8080"}
					m := http.NewServeMux()
					m.Handle("/metrics", promhttp.HandlerFor(
						srv.PrometheusRegistry(),
						promhttp.HandlerOpts{
							EnableOpenMetrics: true,
						},
					))
					httpSrv.Handler = m
					if err := httpSrv.ListenAndServe(); err != nil {
						panic(err)
					}
					return nil
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				srv.GracefulStop()
				return nil
			},
		},
	)
}
