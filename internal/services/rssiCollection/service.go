package rssicollection

import (
	"context"

	//"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	IsExpectedApExisted(ctx context.Context, req *rssiv1.GetCoordinateRequest) (bool, error)
	RegisterNewAp(ctx context.Context, ap models.AP) error
}

type RssiCollectionService struct {
	apCollectionRepo apcollectionrepo.Repository
	cfg              config.ApCollectionServiceConfig
}

func ProvideRssiCollectionService(apCollectionRepo apcollectionrepo.Repository, cfg config.ApCollectionServiceConfig) Service {
	return &RssiCollectionService{
		apCollectionRepo: apCollectionRepo,
		cfg:              cfg,
	}
}

func (s *RssiCollectionService) IsExpectedApExisted(ctx context.Context, req *rssiv1.GetCoordinateRequest) (bool, error) {
	filter := mongodb.Filter{
		"ssid":        req.Signals[0].Ssid,
		"mac_address": req.Signals[0].MacAddress,
	}
	_, err := s.apCollectionRepo.FindOne(ctx, filter)
	if err != nil && err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (s *RssiCollectionService) RegisterNewAp(ctx context.Context, ap models.AP) error {
	if err := s.apCollectionRepo.InsertOne(ctx, ap); err != nil {
		return err
	}
	return nil
}
