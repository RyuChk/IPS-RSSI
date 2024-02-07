package statcollection

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/utils"
	"github.com/rs/zerolog/log"
)

func apToMap(aps *[]models.AP) map[string]string {
	m := make(map[string]string)
	for _, v := range *aps {
		m[v.MacAddress] = v.Name
	}
	return m
}

func rssiInfoToMap(rssi []models.RSSI) map[string]models.RSSI {
	m := make(map[string]models.RSSI)
	for _, v := range rssi {
		m[v.MacAddress] = v
	}

	return m
}

func (s *StatCollectionService) RSSIDataProcessing(ctx context.Context, stat models.RSSIStatModel) ([]models.RSSIDetailStatModel, error) {
	registeredAPs, err := s.GetAllRegisterAPs(ctx)
	if err != nil {
		return []models.RSSIDetailStatModel{}, nil
	}

	filtered := filterAPs(stat.RSSIInfo, registeredAPs)
	stat.RSSIInfo = filtered

	stat.CaculateAverageStrength()
	outlined := filterOutlined(stat.RSSIInfo, s.cfg.OutlineOffset)
	stat.RSSIInfo = outlined

	return s.modelConversion(stat, registeredAPs)
}

func filterOutlined(rssi []models.RSSI, offset float64) []models.RSSI {
	for j, v := range rssi {
		temp := v
		length := len(v.Strength)
		for i := 0; i < length; i++ {
			if temp.Strength[i] >= (v.AverageStrenth+offset) || temp.Strength[i] <= (v.AverageStrenth-offset) {
				temp.Strength = utils.RemoveItemFromSlice(temp.Strength, i)
				temp.CreatedAt = utils.RemoveItemFromSlice(temp.CreatedAt, i)
				i--
				length--
			}
		}
		rssi[j] = temp
	}

	return rssi
}

func filterAPs(rssi []models.RSSI, registeredAPs *[]models.AP) []models.RSSI {
	result := make([]models.RSSI, 0)
	filterMap := rssiInfoToMap(rssi)
	for _, v := range *registeredAPs {
		if data, exist := filterMap[v.MacAddress]; exist {
			result = append(result, data)
		}
	}

	return result
}

func (s *StatCollectionService) modelConversion(stat models.RSSIStatModel, apsList *[]models.AP) ([]models.RSSIDetailStatModel, error) {
	apMap := apToMap(apsList)

	var result []models.RSSIDetailStatModel
	var rssi = stat.RSSIInfo
	timeStampMap := make(map[time.Time][]float64)
	for _, e := range rssi {
		if apName, ok := apMap[strings.ToLower(e.MacAddress)]; ok {
			for i, j := range e.CreatedAt {
				if _, ok := timeStampMap[j]; !ok {
					timeStampMap[j] = makeRSSIArray(apMap, s.cfg.RSSIStrengthDefaultValue)
				}
				if apIndex, exists := autoFindAPIndex(apName); exists {
					timeStampMap[j][apIndex] = e.Strength[i]
				}
			}
		}
	}

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

	return result, nil
}

func makeRSSIArray(apMap map[string]string, rssiStrengthDefaultValue float64) []float64 {
	rssiArray := make([]float64, len(apMap))
	for i := 0; i < len(rssiArray); i++ {
		rssiArray[i] = rssiStrengthDefaultValue
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

	return num - 1, true
}
