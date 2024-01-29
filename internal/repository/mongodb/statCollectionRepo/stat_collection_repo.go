package statcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=stat_collection_repo.go -destination=mock_statCollectionRepo/mock_statcollectionrepo.go -package=mock_statcollectionrepo

const statCollectionName = "signal-stat-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.RSSIStatModel) error
	GetRSSIStats(ctx context.Context) ([]models.RSSIStatModel, error)
	InsertMany(ctx context.Context, documents []models.RSSIDetailStatModel) error
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

func (r *DataCollectionRepo) InsertMany(ctx context.Context, documents []models.RSSIDetailStatModel) error {
	if len(documents) == 0 {
		return nil // No documents to insert
	}

	// Convert RSSIDetailStatModel documents to BSON documents
	var bsonDocuments []interface{}
	for _, doc := range documents {
		bsonDocuments = append(bsonDocuments, doc)
	}

	// Use InsertMany to insert multiple documents
	_, err := r.statCollection.InsertMany(ctx, bsonDocuments)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting documents into the database")
		return err
	}

	log.Debug().Msg("Appended stats to the database")
	return nil
}

func (r *DataCollectionRepo) GetRSSIStats(ctx context.Context) ([]models.RSSIStatModel, error) {
	// Define the filter to retrieve all documents
	filter := bson.M{}

	// Execute the find operation to get all matching records
	cursor, err := r.statCollection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving RSSIStatModels")
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of RSSIStatModel
	var stats []models.RSSIStatModel
	if err := cursor.All(ctx, &stats); err != nil {
		log.Error().Err(err).Msg("Error decoding RSSIStatModels")
		return nil, err
	}

	return stats, nil
}
