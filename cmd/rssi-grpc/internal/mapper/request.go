package mapper

import (
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/constants"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	v1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	sharedv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/shared/rssi/v1"
)

var (
	FromStatCollectionStateEnumToConstant = map[sharedv1.StatCollectionStage]constants.CollectionStage{
		sharedv1.StatCollectionStage_STAT_COLLECTION_STAGE_SINGLE:   constants.CollectionStageSingle,
		sharedv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE: constants.CollectionStageMultiple,
	}
)

func ToRSSIStatModel(req *v1.CollectDataRequest) models.RSSIStatModel {
	output := models.RSSIStatModel{
		RSSIInfo: make([]models.RSSI, len(req.Signals)),
		DeviceInfo: models.DeviceInfo{
			DeviceID: req.DeviceInfo.DeviceId,
			Models:   req.DeviceInfo.Model,
		},
		Position: models.Position{
			X: float64(req.Position.X),
			Y: float64(req.Position.Y),
			Z: float64(req.Position.Z),
		},
		Stage:       FromStatCollectionStateEnumToConstant[req.Stage],
		Duration:    int(req.Duration),
		PollingRate: int(req.PollingRate),
		StartedAt:   req.StartedAt.AsTime(),
		EndedAt:     req.EndedAt.AsTime(),
		CreatedAt:   req.CreatedAt.AsTime(),
	}

	for i, v := range req.Signals {
		temp_time := make([]time.Time, len(v.CreatedAt))

		for j, t := range v.CreatedAt {
			temp_time[j] = t.AsTime()
		}

		output.RSSIInfo[i] = models.RSSI{
			SSID:       v.Ssid,
			MacAddress: v.MacAddress,
			Strength:   v.Strength,
			CreatedAt:  temp_time,
		}
	}

	return output
}

func ToRSSIModel(req *v1.GetCoordinateRequest) []models.RSSI {
	var rssiModel []models.RSSI

	return rssiModel
}
