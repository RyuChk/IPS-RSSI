package di

import (
	grpcProvide "git.cie.com/ips/wire-provider/grpc/provider"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/handler"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/server"
	v1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	server.ProvideGRPCServerCustomizer,
	handler.ProvideStatServer,
	handler.ProvideRssiServer,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handler              *handler.Handlers
	GRPCServerCustomizer grpcProvide.GRPCServerCustomizer
	Stat                 v1.StatCollectionServiceServer
	Rssi                 v1.CoordinateCollectionServiceServer
}
