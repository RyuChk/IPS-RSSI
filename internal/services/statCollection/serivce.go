package statcollection

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
)

type Service interface {
	AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error
	GetSignalStatFromDB(ctx context.Context) error
}

type StatCollectionService struct {
	apCollectionRepo   apcollectionrepo.Repository
	statCollectionRepo statcollectionrepo.Repository
	cfg                config.StatCollectionServiceConfig
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
	// apMap, err := s.apCollectionRepo.GetValidAPsMap(context.Background())
	// if err != nil {

	// }

	//var responseObjects []rssiv1.GetStatDataResponse

	// if err := s.statCollectionRepo.InsertOne(ctx, stat); err != nil {
	// 	return err
	// }
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

func mapRSSIStatModel(stat models.RSSIStatModel, apMap map[string]string) []rssiv1.GetStatDataResponse {
	var result []rssiv1.GetStatDataResponse

	var rssi = stat.RSSIInfo

	maxLen := 0

	for _, ee := range rssi {
		if _, ok := apMap[ee.MacAddress]; ok {
			if len(ee.Strength) > maxLen {
				maxLen = len(ee.Strength)
			}
		}
	}

	rssiArray := makeRSSIArray(maxLen, apMap)

	for i, e := range rssi {
		if apName, ok := apMap[e.MacAddress]; ok {
			if apIndex, exists := findAPIndex(apMap, apName); exists {
				// Update the corresponding value in the rssi array
				if i <= len(e.Strength)-1 {
					rssiArray[i][apIndex] = e.Strength[i]
				}
			}

		}

	}

	// for j,each := range rssi {
	// 	var element rssiv1.GetStatDataResponse
	// 	element.Rssi  = rssiArray[j]
	// 	result = append(result, element)
	// }

	return result
}

func makeRSSIArray(max int, apMap map[string]string) [][]float64 {
	var rssiArray [][]float64
	eachArray := make([]float64, len(apMap))
	for i := range rssiArray {
		eachArray[i] = -99
	}

	for j := 0; j < max; j++ {
		rssiArray = append(rssiArray, eachArray)
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
