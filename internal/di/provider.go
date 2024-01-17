package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/cache"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/minio"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
	rssicollection "github.com/ZecretBone/ips-rssi-service/internal/services/rssiCollection"
	statcollection "github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	statcollectionrepo.ProvideStatCollectionRepo,
	statcollection.ProvideStatCollectionService,
	apcollectionrepo.ProvideApCollectionRepo,
	rssicollection.ProvideRssiCollectionService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideStatCollectionServiceConfig,
	config.ProvideApCollectionServiceConfig,
)

type Locator struct {
	MongoDBConn  wiremongo.Connection
	MinioXConn   wireminio.Connection
	RSSIService  statcollection.Service
	CacheService cache.Service
}
