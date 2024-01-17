package statcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=stat_collection_repo.go -destination=mock_statCollectionRepo/mock_statcollectionrepo.go -package=mock_statcollectionrepo

const statCollectionName = "signal-stat-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.RSSIStatModel) error
}

type DataCollectionRepo struct {
	statCollection *mongo.Collection
}

func ProvideStatCollectionRepo(conn wiremongo.Connection) Repository {
	return &DataCollectionRepo{
		statCollection: conn.Database().Collection(statCollectionName),
	}
}

func (r *DataCollectionRepo) InsertOne(ctx context.Context, document models.RSSIStatModel) error {

	log.Debug().Any("RSSIStatModel", document).Msg("Inserting into db")

	_, err := r.statCollection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	log.Debug().Msg("append stat to server")
	return nil
}
