package models

type AP struct {
	Name       string `bson:"name"`
	SSID       string `bson:"ssid"`
	MacAddress string `bson:"mac_address"`
}
