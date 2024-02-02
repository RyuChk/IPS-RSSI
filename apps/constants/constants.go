package constants

type CollectionStage string

var (
	CollectionStageSingle   CollectionStage = "SINGLE"   //Testing RSSI reciever stage
	CollectionStageMultiple CollectionStage = "MULTIPLE" //All aps collection
)

type CachePrefix string

var (
	GetAllAPsCachePrefix CachePrefix = "FIND:AP"
)
