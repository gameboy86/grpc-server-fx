package grpcserverfx

import "google.golang.org/grpc"

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
