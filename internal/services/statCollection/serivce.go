package statcollection

import (
	"context"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
)

type Service interface {
	AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error
}

type StatCollectionService struct {
	statCollectionRepo statcollectionrepo.Repository
	cfg                config.StatCollectionServiceConfig
}

func ProvideStatCollectionService(statCollectionRepo statcollectionrepo.Repository, cfg config.StatCollectionServiceConfig) Service {
	return &StatCollectionService{
		statCollectionRepo: statCollectionRepo,
		cfg:                cfg,
	}
}

func (s *StatCollectionService) AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error {
	stat.CaculateAverageStrength()
	if err := s.statCollectionRepo.InsertOne(ctx, stat); err != nil {
		return err
	}
	return nil
}
