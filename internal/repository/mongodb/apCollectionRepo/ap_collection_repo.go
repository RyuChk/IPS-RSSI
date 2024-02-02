package apcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	//"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=ap_collection_repo.go -destination=mock_apCollectionRepo/mock_apcollectionrepo.go -package=mock_apcollectionrepo

const apCollectionName = "valid-ap-collection"

type Repository interface {
	InsertOne(ctx context.Context, request models.AP) error
	FindOne(ctx context.Context, filter mongodb.Filter) (models.AP, error)
	Find(ctx context.Context, filter mongodb.Filter) (*[]models.AP, error)
}

type ApCollectionRepo struct {
	apCollection *mongo.Collection
}

func ProvideApCollectionRepo(conn wiremongo.Connection) Repository {
	return &ApCollectionRepo{
		apCollection: conn.Database().Collection(apCollectionName),
	}
}

func (r *ApCollectionRepo) InsertOne(ctx context.Context, AP models.AP) error {
	_, err := r.apCollection.InsertOne(ctx, AP)
	if err != nil {
		return err
	}
	log.Debug().Msg("append ap to server")
	return nil
}

func (r *ApCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.AP, error) {
	result := models.AP{}
	err := r.apCollection.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, nil
	}

	return result, nil
}

func (r *ApCollectionRepo) Find(ctx context.Context, filter mongodb.Filter) (*[]models.AP, error) {
	result := &[]models.AP{}

	curr, err := r.apCollection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	if err := curr.All(ctx, result); err != nil {
		return nil, err
	}
	return result, nil
}
