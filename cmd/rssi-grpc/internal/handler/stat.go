package handler

import (
	"context"
	"net/http"

	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/mapper"
	"github.com/ZecretBone/ips-rssi-service/internal/errorx"
	statv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	statcollection "github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

type StatV1Impl struct {
	statCollectionService statcollection.Service
	statv1.UnimplementedStatCollectionServiceServer
}

func ProvideStatServer(statCollectionService statcollection.Service) statv1.StatCollectionServiceServer {
	return &StatV1Impl{
		statCollectionService: statCollectionService,
	}
}

func (s *StatV1Impl) CollectData(ctx context.Context, req *statv1.CollectDataRequest) (*statv1.CollectDataResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errorx.NewAPIError(http.StatusBadRequest, "no header were found")
	}
	log.Debug().Any("header", md.Get("x-device-id"))

	if err := s.statCollectionService.AddSignalStatToDB(ctx, mapper.ToRSSIStatModel(req)); err != nil {
		return nil, errorx.NewAPIError(http.StatusBadRequest, err.Error())
	}
	return &statv1.CollectDataResponse{}, nil
}

func (s *StatV1Impl) GetStatData(ctx context.Context, req *statv1.GetStatDataRequest) (*statv1.GetStatDataResponse, error) {

	if err := s.statCollectionService.GetSignalStatFromDB(ctx); err != nil {
		return nil, errorx.NewAPIError(http.StatusBadRequest, err.Error())
	}
	return &statv1.GetStatDataResponse{}, nil
}
