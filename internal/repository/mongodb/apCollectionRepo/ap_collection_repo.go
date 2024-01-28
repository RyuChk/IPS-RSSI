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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=ap_collection_repo.go -destination=mock_apCollectionRepo/mock_apcollectionrepo.go -package=mock_apcollectionrepo

const apCollectionName = "valid-ap-collection"

type Repository interface {
	InsertOne(ctx context.Context, request *rssiv1.RegisterApRequest) error
	IsExpectedApExisted(ctx context.Context, request *rssiv1.GetCoordinateRequest) (bool, error)
	GetValidAPsMap(ctx context.Context) (map[string]string, error)
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

func (r *ApCollectionRepo) GetValidAPsMap(ctx context.Context) (map[string]string, error) {
	// Build the filter to match names starting with "AP"
	filter := bson.M{
		"name": bson.M{
			"$regex": "^AP",
		},
	}

	// Specify the sorting criteria
	sort := bson.D{
		{"name", 1}, // Sort by name in ascending order
	}

	// Execute the find operation to get matching records and sort them
	cursor, err := r.apCollection.Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving valid APs")
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of YourAPStruct
	var aps []rssiv1.RegisterApRequest
	if err := cursor.All(ctx, &aps); err != nil {
		log.Error().Err(err).Msg("Error decoding APs")
		return nil, err
	}

	// Create a map to store the result
	resultMap := make(map[string]string)

	// Populate the map with mac_address as key and AP name as value
	// for _, ap := range aps {
	// 	resultMap[ap.MacAddress] = ap.Name
	// }

	return resultMap, nil
}
