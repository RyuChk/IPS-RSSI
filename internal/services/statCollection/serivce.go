package statcollection

import (
	"context"
	"fmt"
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/constants"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/ZecretBone/ips-rssi-service/utils/converter"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/cache"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
	trainstatcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/trainstatCollectionRepo"
)

type Service interface {
	AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error
	GetSignalStatFromDB(ctx context.Context) error
	DoDataProcessingFromTimeStamp(ctx context.Context, start time.Time, end time.Time) (models.ReDoDataProcessingResult, error)
}

type StatCollectionService struct {
	localCache              cache.Service
	apCollectionRepo        apcollectionrepo.Repository
	statCollectionRepo      statcollectionrepo.Repository
	trainstatCollectionRepo trainstatcollectionrepo.Repository
	cfg                     config.StatCollectionServiceConfig
}

func ProvideStatCollectionService(localCache cache.Service, apCollectionRepo apcollectionrepo.Repository, statCollectionRepo statcollectionrepo.Repository, trainstatCollectionRepo trainstatcollectionrepo.Repository, cfg config.StatCollectionServiceConfig) Service {
	return &StatCollectionService{
		localCache:              localCache,
		apCollectionRepo:        apCollectionRepo,
		statCollectionRepo:      statCollectionRepo,
		trainstatCollectionRepo: trainstatCollectionRepo,
		cfg:                     cfg,
	}
}

func (s *StatCollectionService) AddSignalStatToDB(ctx context.Context, stats models.RSSIStatModel) error {
	rssiNewModel, err := s.RSSIDataProcessing(ctx, stats)
	if err != nil {
		return err
	}

	if err := s.trainstatCollectionRepo.InsertMany(ctx, rssiNewModel); err != nil {
		return err
	}

	if err := s.statCollectionRepo.InsertOne(ctx, stats); err != nil {
		return err
	}

	return nil
}

func (s *StatCollectionService) GetAllRegisterAPs(ctx context.Context) (*[]models.AP, error) {
	result := make([]models.AP, 0)

	cache, err := s.localCache.Get(string(constants.GetAllAPsCachePrefix))
	if err != nil {
		APs, err := s.apCollectionRepo.Find(ctx, mongodb.Filter{})
		if err != nil {
			return nil, err
		}

		bytes, err := converter.ToByte(APs)
		if err != nil {
			return nil, err
		}

		if err := s.localCache.Set(string(constants.GetAllAPsCachePrefix), bytes); err != nil {
			return nil, err
		}

		return APs, nil
	}

	if err := converter.FromByte(cache, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *StatCollectionService) GetSignalStatFromDB(ctx context.Context) error {
	allStat, err := s.statCollectionRepo.GetRSSIStats(ctx)
	if err != nil {
		return err
	}

	registeredAPs, err := s.GetAllRegisterAPs(ctx)
	if err != nil {
		return err
	}

	for _, stat := range allStat {
		output := processRSSIStatModel(stat, apToMap(registeredAPs))
		fmt.Printf("%+v\n", output)
	}

	return nil
}

func (s *StatCollectionService) DoDataProcessingFromTimeStamp(ctx context.Context, start time.Time, end time.Time) (models.ReDoDataProcessingResult, error) {
	log.Info().Str("start_at", start.String()).Str("end_at", end.String()).Msg("process time spand")
	result := models.ReDoDataProcessingResult{
		ErrorData: make([]models.ReDoError, 0),
	}

	filter, _ := mongodb.AddFilter(
		mongodb.Filter{"created_at": mongodb.Filter{"$gte": start}},
		mongodb.Filter{"created_at": mongodb.Filter{"$lte": end}},
	)

	DataToProcess, err := s.statCollectionRepo.Find(ctx, filter)
	if err != nil {
		return models.ReDoDataProcessingResult{}, status.Error(codes.Internal, err.Error())
	}

	for _, v := range DataToProcess {
		processedData, err := s.RSSIDataProcessing(ctx, v)
		if err != nil {
			result.TotalError++
			result.ErrorData = append(result.ErrorData, models.ReDoError{
				Data:    v,
				Message: err.Error(),
			})
			continue
		}

		if err := s.trainstatCollectionRepo.InsertMany(ctx, processedData); err != nil {
			result.TotalError++
			result.ErrorData = append(result.ErrorData, models.ReDoError{
				Data:    v,
				Message: err.Error(),
			})
			continue
		}

		result.TotalDataProcessed++
		result.TotalRowAdded += len(processedData)
	}

	log.Debug().Any("result_struct", result).Msg("finish process data")
	return result, nil
}
