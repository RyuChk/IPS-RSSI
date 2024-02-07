package handler

import (
	"context"
	"net/http"

	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/mapper"
	"github.com/ZecretBone/ips-rssi-service/internal/errorx"
	statv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	statcollection "github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
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

func (s *StatV1Impl) ReDoDataProcessing(ctx context.Context, req *statv1.ReDoDataProcessingRequest) (*statv1.ReDoDataProcessingResponse, error) {
	result, err := s.statCollectionService.DoDataProcessingFromTimeStamp(ctx, req.StartAt.AsTime(), req.EndAt.AsTime())
	if err != nil {
		return nil, err
	}

	return &statv1.ReDoDataProcessingResponse{
		TotalDataProcessed: int32(result.TotalDataProcessed),
		TotalRowAdded:      int32(result.TotalRowAdded),
	}, nil
}
