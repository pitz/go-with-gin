package fixtures

import (
	"github.com/stretchr/testify/mock"
	"pitzdev/web-service-gin/models"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetScore(analyse *models.Analyse, ch chan models.Score) {
	args := m.Called(analyse, ch)
	score := args.Get(0).(models.Score)

	ch <- score
}
