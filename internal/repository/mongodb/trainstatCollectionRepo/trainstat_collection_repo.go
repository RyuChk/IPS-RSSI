package trainstatcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/rs/zerolog/log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=trainstat_collection_repo.go -destination=mock_trainstatCollectionRepo/mock_trainstatcollectionrepo.go -package=mock_trainstatcollectionrepo

const trainstatCollectionName = "signal-trainstat-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.RSSIStatModel) error
	InsertMany(ctx context.Context, documents []models.RSSIDetailStatModel) error
}

type TrainDataCollectionRepo struct {
	trainstatCollection *mongo.Collection
}

func ProvideTrainStatCollectionRepo(conn wiremongo.Connection) Repository {
	return &TrainDataCollectionRepo{
		trainstatCollection: conn.Database().Collection(trainstatCollectionName),
	}
}

func (r *TrainDataCollectionRepo) InsertOne(ctx context.Context, document models.RSSIStatModel) error {

	log.Debug().Any("RSSIStatModel", document).Msg("Inserting into db")

	_, err := r.trainstatCollection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	log.Debug().Msg("append stat to server")
	return nil
}

func (r *TrainDataCollectionRepo) InsertMany(ctx context.Context, documents []models.RSSIDetailStatModel) error {
	if len(documents) == 0 {
		return nil // No documents to insert
	}

	// Convert RSSIDetailStatModel documents to BSON documents
	var bsonDocuments []interface{}
	for _, doc := range documents {
		bsonDocuments = append(bsonDocuments, doc)
	}

	// Use InsertMany to insert multiple documents
	_, err := r.trainstatCollection.InsertMany(ctx, bsonDocuments)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting documents into the database")
		return err
	}

	log.Debug().Msg("Appended stats to the database")
	return nil
}
