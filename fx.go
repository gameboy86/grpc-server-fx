package grpcserverfx

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

func NewPrometheus() *PrometheusGRPC {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets(
				[]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120},
			),
		),
	)
	return &PrometheusGRPC{
		Registry:      prometheus.NewRegistry(),
		ServerMetrics: srvMetrics,
	}
}

var Module = fx.Module(
	"grpc_server",
	fx.Invoke(
		func() *PrometheusGRPC {
			if true {
				return NewPrometheus()
			}
			return nil
		},
	),
	fx.Provide(
		fx.Annotate(
			NewGRPCServer,
			fx.As(new(GRPCServerer)),
		),
		fx.Annotate(
			NewGRPCPromServerMetrics,
			fx.As(new(PrometheusServerMetrics)),
		),
		NewListener,
		// NewPrometheus,
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
