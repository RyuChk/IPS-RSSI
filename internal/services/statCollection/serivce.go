package statcollection

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"

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

func ProvideStatCollectionService(apCollectionRepo apcollectionrepo.Repository, statCollectionRepo statcollectionrepo.Repository, cfg config.StatCollectionServiceConfig) Service {
	return &StatCollectionService{
		apCollectionRepo:   apCollectionRepo,
		statCollectionRepo: statCollectionRepo,
		cfg:                cfg,
	}
}

func (s *StatCollectionService) AddSignalStatToDB(ctx context.Context, stats models.RSSIStatModel) error {
	stats.CaculateAverageStrength()

	//convert to new struture
	apMap, err := s.apCollectionRepo.GetValidAPsMap(context.Background())
	if err != nil {

	}

	rssiNewModel := mapRSSIStatModel(stats, apMap)

	if err := s.trainstatCollectionRepo.InsertMany(ctx, rssiNewModel); err != nil {
		return err
	}

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
	var result []models.RSSIDetailStatModel

	var rssi = stat.RSSIInfo
	timeStampMap := make(map[time.Time][]float64)

	for _, e := range rssi {
		if apName, ok := apMap[e.MacAddress]; ok {
			for i, j := range e.CreatedAt {
				if _, ok := timeStampMap[j]; !ok {
					timeStampMap[j] = makeRSSIArray(apMap)
				}
				if apIndex, exists := findAPIndex(apMap, apName); exists {
					timeStampMap[j][apIndex] = e.Strength[i]
				}
			}
		}
	}

	for j, eachRSSI := range timeStampMap {
		var element models.RSSIDetailStatModel
		element.RSSI = eachRSSI
		element.Model = stat.DeviceInfo.Models
		element.PollingRate = stat.PollingRate
		element.Stage = stat.Stage
		element.CreatedAt = j
		result = append(result, element)
	}

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
