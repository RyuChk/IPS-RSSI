package statcollection

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/rs/zerolog/log"

	//rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
	trainstatcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/trainstatCollectionRepo"
)

type Service interface {
	AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error
	GetSignalStatFromDB(ctx context.Context) error
}

type StatCollectionService struct {
	apCollectionRepo        apcollectionrepo.Repository
	statCollectionRepo      statcollectionrepo.Repository
	trainstatCollectionRepo trainstatcollectionrepo.Repository
	cfg                     config.StatCollectionServiceConfig
}

func ProvideStatCollectionService(apCollectionRepo apcollectionrepo.Repository, statCollectionRepo statcollectionrepo.Repository, trainstatCollectionRepo trainstatcollectionrepo.Repository, cfg config.StatCollectionServiceConfig) Service {
	return &StatCollectionService{
		apCollectionRepo:        apCollectionRepo,
		statCollectionRepo:      statCollectionRepo,
		trainstatCollectionRepo: trainstatCollectionRepo,
		cfg:                     cfg,
	}
}

func (s *StatCollectionService) AddSignalStatToDB(ctx context.Context, stats models.RSSIStatModel) error {
	stats.CaculateAverageStrength()

	//convert to new struture
	apMap, err := s.apCollectionRepo.GetValidAPsMap(context.Background())
	if err != nil {
		log.Debug().Msg("show err getting ap map")
	}

	log.Debug().Msg("show1")

	rssiNewModel := mapRSSIStatModel(stats, apMap)

	log.Debug().Msg("show imany")

	if err := s.trainstatCollectionRepo.InsertMany(ctx, rssiNewModel); err != nil {
		return err
	}
	log.Debug().Msg("show ione")

	if err := s.statCollectionRepo.InsertOne(ctx, stats); err != nil {
		return err
	}

	return nil
}

func (s *StatCollectionService) GetSignalStatFromDB(ctx context.Context) error {
	allStat, err := s.statCollectionRepo.GetRSSIStats(ctx)
	if err != nil {
		return err
	}

	apMap, err := s.apCollectionRepo.GetValidAPsMap(context.Background())
	if err != nil {

	}

	for _, stat := range allStat {
		output := processRSSIStatModel(stat, apMap)
		fmt.Printf("%+v\n", output)
	}

	return nil
}

func mapRSSIStatModel(stat models.RSSIStatModel, apMap map[string]string) []models.RSSIDetailStatModel {
	log.Debug().Msg("show ver2")
	log.Debug().Msg("show1.5")
	var result []models.RSSIDetailStatModel
	logMap := fmt.Sprintf("show Map as string: %v", apMap)
	log.Debug().Msg(logMap)
	var rssi = stat.RSSIInfo
	timeStampMap := make(map[time.Time][]float64)
	log.Debug().Msg("show2")

	logRSSI := fmt.Sprintf("show RSSI as string: %v", rssi)
	log.Debug().Msg(logRSSI)
	for _, e := range rssi {
		log.Debug().Msg("show looping rssi")
		logMac := fmt.Sprintf("show MAC from rssi as string: %v", e.MacAddress)
		log.Debug().Msg(logMac)
		if apName, ok := apMap[strings.ToLower(e.MacAddress)]; ok {
			log.Debug().Msg("show adding matched ap")
			for i, j := range e.CreatedAt {
				if _, ok := timeStampMap[j]; !ok {
					log.Debug().Msg("show adding new time stamp")
					timeStampMap[j] = makeRSSIArray(apMap)
				}
				// if apIndex, exists := findAPIndex(apMap, apName); exists {
				// 	logIndex := fmt.Sprintf("show AP Index as string: %v", apIndex)
				// 	log.Debug().Msg(logIndex)
				// 	log.Debug().Msg("show adding new existed time stamp")
				// 	timeStampMap[j][apIndex] = e.Strength[i]
				// }
				if apIndex, exists := autoFindAPIndex(apName); exists {
					logIndex := fmt.Sprintf("show AP Index as string: %v", apIndex)
					log.Debug().Msg(logIndex)
					log.Debug().Msg("show adding new existed time stamp")
					timeStampMap[j][apIndex] = e.Strength[i]
				}
			}
		}
	}
	log.Debug().Msg("show2.5")
	logStampMap := fmt.Sprintf("show FULL timestamp map as string: %v", timeStampMap)
	log.Debug().Msg(logStampMap)

	for j, eachRSSI := range timeStampMap {
		log.Debug().Msg("show2.7 adding rssi timeStamp")
		var element models.RSSIDetailStatModel
		element.RSSI = eachRSSI
		element.X = float32(stat.Position.X)
		element.Y = float32(stat.Position.Y)
		element.Z = float32(stat.Position.Z)
		element.Model = stat.DeviceInfo.Models
		element.PollingRate = stat.PollingRate
		element.Stage = stat.Stage
		element.CreatedAt = j
		result = append(result, element)
	}
	log.Debug().Msg("show3")

	return result
}

// func makeRSSIArray(max int, apMap map[string]string) [][]float64 {
// 	var rssiArray [][]float64
// 	eachArray := make([]float64, len(apMap))
// 	for i := range rssiArray {
// 		eachArray[i] = -99
// 	}

// 	for j := 0; j < max; j++ {
// 		rssiArray = append(rssiArray, eachArray)
// 	}

// 	return rssiArray
// }

func makeRSSIArray(apMap map[string]string) []float64 {
	rssiArray := make([]float64, len(apMap))
	for i := 0; i < len(rssiArray); i++ {
		rssiArray[i] = -99
	}
	return rssiArray
}

// processRSSIStatModel creates the desired output for a given RSSIStatModel
func processRSSIStatModel(stat models.RSSIStatModel, apMap map[string]string) map[string]interface{} {
	result := make(map[string]interface{})

	// Initialize the rssi array with -99 for unmatched APs
	rssiArray := make([]float64, len(apMap))
	for i := range rssiArray {
		rssiArray[i] = -99
	}

	// Map each mac_address to AP
	for _, rssiInfo := range stat.RSSIInfo {
		if apName, ok := apMap[rssiInfo.MacAddress]; ok {
			// Find the index of the AP in the map
			if apIndex, exists := findAPIndex(apMap, apName); exists {
				// Update the corresponding value in the rssi array
				rssiArray[apIndex] = rssiInfo.Strength[0] // Assuming only the first strength value is used
			}
		}
	}

	// Populate the result map
	result["rssi"] = rssiArray
	result["model"] = stat.DeviceInfo.Models
	result["x"] = stat.Position.X
	result["y"] = stat.Position.Y
	result["z"] = stat.Position.Z
	result["poll_rate"] = stat.PollingRate
	//result["collection_stage"] = stat.CollectionStage

	return result
}

// findAPIndex finds the index of an AP in the AP map
func findAPIndex(apMap map[string]string, apName string) (int, bool) {
	for i, name := range apMap {
		if name == apName {
			// Convert i to int before returning
			index, err := strconv.Atoi(i)
			if err != nil {
				return -1, false
			}
			return index, true
		}
	}
	return -1, false
}

func autoFindAPIndex(apName string) (int, bool) {
	numStr := ""
	for _, char := range apName {
		if char >= '0' && char <= '9' {
			numStr += string(char)
		}
	}

	// Convert the numeric string to an integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1, false
	}

	return num, true
}
