package server

import (
	wiregrpc "git.cie.com/ips/wire-provider/grpc"
	"git.cie.com/ips/wire-provider/grpc/provider"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/handler"
)

type grpcServerCustomizer struct {
	handlers *handler.Handlers
}

func (gc *grpcServerCustomizer) Register(server wiregrpc.Server) error {
	handler.RegisterService(server, gc.handlers)
	return nil
}

func (gc *grpcServerCustomizer) Configrue(builder wiregrpc.Builder) error {
	builder.WithServiceName("bff-service").WithListenAddr(":6000")
	return nil
}

func ProvideGRPCServerCustomizer(handlers *handler.Handlers) provider.GRPCServerCustomizer {
	return &grpcServerCustomizer{
		handlers: handlers,
	}
}
