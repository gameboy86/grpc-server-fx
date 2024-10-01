package grpcserverfx

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

type PrometheusGRPC struct {
	Registry      *prometheus.Registry
	ServerMetrics *grpcprom.ServerMetrics
}

func (p *PrometheusGRPC) InitializeMetrics(srv *grpc.Server) {
	p.Registry.MustRegister(p.ServerMetrics)
	p.ServerMetrics.InitializeMetrics(srv)
}

func NewPrometheusGRPC(
	registry *prometheus.Registry,
	serverMetrics *grpcprom.ServerMetrics,
) *PrometheusGRPC {
	return &PrometheusGRPC{
		Registry:      registry,
		ServerMetrics: serverMetrics,
	}
}
