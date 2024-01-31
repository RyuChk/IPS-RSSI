package models

import "time"

type RSSI struct {
	SSID           string      `bson:"ssid"`             //APs SSID
	MacAddress     string      `bson:"mac_address"`      // MAC Address of APs
	Strength       []float64   `bson:"strength"`         //RSSI signal strength in Dbm.
	AverageStrenth float64     `bson:"average_strength"` //Average of all rssi strength
	CreatedAt      []time.Time `bson:"created_at"`       //RSSI signal capture time
}

type Position struct {
	X float64 `bson:"x"` // X position in meter
	Y float64 `bson:"y"` // Y position in meter
	Z float64 `bson:"z"` // Z position in floor
}

type DeviceInfo struct {
	DeviceID string `bson:"device_id"` //Device ID
	Models   string `bson:"models"`    //Device models eg.Samsung
}

type AccessPoint struct {
	Name       string `bson:"name"`
	SSID       string `bson:"ssid"`        //APs SSID
	MacAddress string `bson:"mac_address"` // MAC Address of APs
}

// type RSSICoorRequest struct{
// 	Signals []RSSI `bson:"signals"`
// 	DeviceInfo DeviceInfo
// }
