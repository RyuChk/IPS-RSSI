package handler

import (
	"context"
	"net/http"

	//"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/mapper"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/mapper"
	"github.com/ZecretBone/ips-rssi-service/internal/errorx"
	rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	shared_rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/shared/rssi/v1"

	//statv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	rssicollection "github.com/ZecretBone/ips-rssi-service/internal/services/rssiCollection"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

type RssiV1Impl struct {
	rssiCollectionService rssicollection.Service
	rssiv1.UnimplementedCoordinateCollectionServiceServer
}

func ProvideRssiServer(rssiCollectionService rssicollection.Service) rssiv1.CoordinateCollectionServiceServer {
	return &RssiV1Impl{
		rssiCollectionService: rssiCollectionService,
	}
}

func (s *RssiV1Impl) GetCoordinate(ctx context.Context, req *rssiv1.GetCoordinateRequest) (*rssiv1.GetCoordinateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errorx.NewAPIError(http.StatusBadRequest, "no header were found")
	}
	log.Debug().Any("header", md.Get("x-device-id"))

	// valid, err := s.rssiCollectionService.IsExpectedApExisted(ctx, mapper.ToRSSIModel(req))
	valid, err := s.rssiCollectionService.IsExpectedApExisted(ctx, req)
	if err != nil {
		return nil, errorx.NewAPIError(http.StatusBadRequest, err.Error())
	}
	if !valid {
		return nil, errorx.NewAPIError(http.StatusBadRequest, "no valid APs were found")
	}

	// TODO: get coor by req

	//dummy resp
	resp := rssiv1.GetCoordinateResponse{
		Position: &shared_rssiv1.Position{
			X: 1,
			Y: 2,
			Z: 10,
		},
	}

	return &resp, nil
}

func (s *RssiV1Impl) RegisterAp(ctx context.Context, req *rssiv1.RegisterApRequest) (*rssiv1.RegisterApResponse, error) {
	if err := s.rssiCollectionService.RegisterNewAp(ctx, mapper.ToAPModel(req)); err != nil {
		return nil, errorx.NewAPIError(http.StatusBadRequest, err.Error())
	}
	return &rssiv1.RegisterApResponse{}, nil
}
