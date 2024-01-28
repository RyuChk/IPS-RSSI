package statcollection_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	mock_apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo/mock_apCollectionRepo"
	mock_statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo/mock_statCollectionRepo"
	statcollection "github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type StatCollectionServiceTestSuite struct {
	suite.Suite

	ctrl                   *gomock.Controller
	mockStatCollectionRepo *mock_statcollectionrepo.MockRepository
	mockAPCollectionRepo   *mock_apcollectionrepo.MockRepository // Add this line

	statCollectionService statcollection.Service
	timeLayout            string
	ctx                   context.Context
}

func TestStatCollectionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StatCollectionServiceTestSuite))
}

func (s *StatCollectionServiceTestSuite) SetupSuite() {
	s.Init()
}

func (s *StatCollectionServiceTestSuite) Init() {
	s.ctrl = gomock.NewController(s.T())
	s.timeLayout = "2006-01-02 15:04:05"
	s.mockStatCollectionRepo = mock_statcollectionrepo.NewMockRepository(s.ctrl)

	//s.statCollectionService = statcollection.ProvideStatCollectionService(s.mockStatCollectionRepo, config.ProvideStatCollectionServiceConfig())
	// Assuming that MockRepository implements statcollectionrepo.Repository
	s.statCollectionService = statcollection.ProvideStatCollectionService(
		s.mockAPCollectionRepo,
		s.mockStatCollectionRepo, // Assuming s.mockAPCollectionRepo is of type apcollectionrepo.Repository
		config.ProvideStatCollectionServiceConfig(),
	)

}

// func (s *StatCollectionServiceTestSuite) SetupTest() {
// 	s.ctrl = gomock.NewController(s.T())
// 	s.mockStatCollectionRepo = mock_statcollectionrepo.NewMockRepository(s.ctrl)
// 	s.mockAPCollectionRepo = mock_apcollectionrepo.NewMockRepository(s.ctrl)

// 	// Initialize other fields if needed
// 	s.statCollectionService = statcollection.ProvideStatCollectionService(
// 		s.mockAPCollectionRepo,
// 		s.mockStatCollectionRepo,
// 		config.ProvideStatCollectionServiceConfig(),
// 	)

// 	s.timeLayout = "2006-01-02T15:04:05.000Z"
// 	s.ctx = context.TODO()
// }

// func (s *StatCollectionServiceTestSuite) TearDownTest() {
// 	s.ctrl.Finish()
// }

func initRSSIModel() models.RSSIStatModel {
	currentTime, _ := time.Parse("2006-01-02 15:04:05", "2023-12-29 10:10:0")

	return models.RSSIStatModel{
		RSSIInfo: []models.RSSI{
			{
				SSID:      "AP-1",
				Strength:  []float64{60.3},
				CreatedAt: []time.Time{currentTime.Add(time.Second)},
			},
		},
		DeviceInfo: models.DeviceInfo{
			DeviceID: "<device-id>",
			Models:   "<device-model>",
		},
		Duration:  10,
		StartedAt: currentTime,
		EndedAt:   currentTime.Add(time.Second * 10),
		CreatedAt: currentTime.Add(time.Second * 12),
	}
}

func (s *StatCollectionServiceTestSuite) Test_Add_Signal_Stat_To_DB_SUCCESS() {
	RSSIModel := initRSSIModel()

	s.mockStatCollectionRepo.EXPECT().InsertOne(gomock.Any(), RSSIModel).Return(nil)

	err := s.statCollectionService.AddSignalStatToDB(s.ctx, RSSIModel)
	s.NoError(err)
}

func (s *StatCollectionServiceTestSuite) Test_Add_Signal_Stat_To_DB_FAILED() {
	RSSIModel := initRSSIModel()

	s.mockStatCollectionRepo.EXPECT().InsertOne(gomock.Any(), RSSIModel).Return(errors.New("some error"))

	err := s.statCollectionService.AddSignalStatToDB(s.ctx, RSSIModel)
	s.Error(err, "some error")
}
