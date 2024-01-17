//go:build wireinject
// +build wireinject

package di

import (
	"git.cie.com/ips/wire-provider/grpc/provider"
	internalDi "github.com/ZecretBone/ips-rssi-service/internal/di"
	"github.com/google/wire"
)

var BaseBindingSet = wire.NewSet(
	internalDi.DatabaseSet,
	internalDi.ConfigSet,
	internalDi.ProviderSet,
)

var MainBindingSet = wire.NewSet(
	ProviderSet,
	BaseBindingSet,
	provider.WireSet,
	wire.Struct(new(Container), "*"),
)

func InitializeContainer() (*Container, func(), error) {
	wire.Build(MainBindingSet)
	return &Container{}, func() {}, nil
}
