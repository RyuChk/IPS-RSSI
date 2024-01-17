package apcollectionrepo

import (
	"context"
	"errors"
	"fmt"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	//"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=ap_collection_repo.go -destination=mock_apCollectionRepo/mock_apcollectionrepo.go -package=mock_apcollectionrepo

const apCollectionName = "valid-ap-collection"

type Repository interface {
	InsertOne(ctx context.Context, request *rssiv1.RegisterApRequest) error
	IsExpectedApExisted(ctx context.Context, request *rssiv1.GetCoordinateRequest) (bool, error)
}

type ApCollectionRepo struct {
	apCollection *mongo.Collection
}

func ProvideApCollectionRepo(conn wiremongo.Connection) Repository {
	return &ApCollectionRepo{
		apCollection: conn.Database().Collection(apCollectionName),
	}
}

func (r *ApCollectionRepo) InsertOne(ctx context.Context, request *rssiv1.RegisterApRequest) error {

	bson := bson.M{
		"ssid":        request.Ssid,
		"mac_address": request.MacAddress,
	}
	fmt.Println("mybson: ")
	fmt.Println(bson)
	log.Debug().Any("RSSIApModel", bson).Msg("Inserting into db")

	_, err := r.apCollection.InsertOne(ctx, bson)
	if err != nil {
		return err
	}
	log.Debug().Msg("append ap to server")
	return nil
}

// TODO add get valid-ap
func (r *ApCollectionRepo) IsExpectedApExisted(ctx context.Context, request *rssiv1.GetCoordinateRequest) (bool, error) {
	// Build the filter based on the SSID and MacAddress in the request
	filter := bson.M{
		"ssid":        request.Signals[0].Ssid,
		"mac_address": request.Signals[0].MacAddress,
	}

	// Execute the find operation to check if a matching record exists
	result := r.apCollection.FindOne(ctx, filter)

	// Check for errors
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// No matching document found
			return false, nil
		}
		// Other error occurred
		log.Error().Err(result.Err()).Msg("Error checking for existing AP")
		return false, result.Err()
	}

	// Matching document found
	return true, nil
}
