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

func (s *RSSIStatModel) CaculateAverageStrength() {
	for i, v := range s.RSSIInfo {
		var t float64
		for _, v := range v.Strength {
			t += v
		}

		s.RSSIInfo[i].AverageStrenth = float64(t) / float64(len(v.Strength))
	}
}
