package models

import (
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/constants"
)

type RSSIStatModel struct {
	RSSIInfo    []RSSI                    `bson:"rssi_info"`
	DeviceInfo  DeviceInfo                `bson:"device_info"`
	Stage       constants.CollectionStage `bson:"collection_stage"`
	Position    Position                  `bson:"position"`
	Duration    int                       `bson:"duration"`
	PollingRate int                       `bson:"polling_rate"`
	StartedAt   time.Time                 `bson:"started_at"`
	EndedAt     time.Time                 `bson:"ended_at"`
	CreatedAt   time.Time                 `bson:"created_at"`
}

type RSSIDetailStatModel struct {
	RSSI        []float64                 `bson:"rssi"`
	Model       string                    `bson:"model"`
	X           float32                   `bson:"x"`
	Y           float32                   `bson:"y"`
	Z           float32                   `bson:"z"`
	CreatedAt   time.Time                 `bson:"created_at"`
	PollingRate int                       `bson:"polling_rate"`
	Stage       constants.CollectionStage `bson:"collection_stage"`
}

type ReDoDataProcessingResult struct {
	TotalDataProcessed int
	TotalRowAdded      int
	TotalError         int
	ErrorData          []ReDoError
}

type ReDoError struct {
	Data    RSSIStatModel
	Message string
}

func (s *RSSIStatModel) CaculateAverageStrength() {
	for i, v := range s.RSSIInfo {
		var t float64
		for _, v := range v.Strength {
			t += v
		}

		s.RSSIInfo[i].AverageStrenth = float64(t) / float64(len(v.Strength))
	}
}
