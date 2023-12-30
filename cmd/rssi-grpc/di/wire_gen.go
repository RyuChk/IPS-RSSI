// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/RyuChk/wire-provider/grpc/provider"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/handler"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/server"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/ZecretBone/ips-rssi-service/internal/di"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
	"github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
	"github.com/google/wire"
)

// Injectors from di.go:

func InitializeContainer() (*Container, func(), error) {
	mongodbConfig := config.ProvideMongoxConfig()
	connection, cleanup, err := mongodb.ProvideMongoDBService(mongodbConfig)
	if err != nil {
		return nil, nil, err
	}
	repository := statcollectionrepo.ProvideStatCollectionRepo(connection)
	statCollectionServiceConfig := config.ProvideStatCollectionServiceConfig()
	service := statcollection.ProvideStatCollectionService(repository, statCollectionServiceConfig)
	statCollectionServiceServer := handler.ProvideStatServer(service)
	handlers := &handler.Handlers{
		Stat: statCollectionServiceServer,
	}
	grpcServerCustomizer := server.ProvideGRPCServerCustomizer(handlers)
	grpcServer, cleanup2, err := provider.ProvideGRPCServer(grpcServerCustomizer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	container := &Container{
		server: grpcServer,
	}
	return container, func() {
		cleanup2()
		cleanup()
	}, nil
}

// di.go:

var BaseBindingSet = wire.NewSet(di.DatabaseSet, di.ConfigSet, di.ProviderSet)

var MainBindingSet = wire.NewSet(
	ProviderSet,
	BaseBindingSet, provider.WireSet, wire.Struct(new(Container), "*"),
)
