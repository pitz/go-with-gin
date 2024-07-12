package fixtures

import (
	"github.com/stretchr/testify/mock"
	"pitzdev/web-service-gin/models"
)

type MockAnalyseController struct {
	mock.Mock
}

func (m *MockAnalyseController) ScheduleExecution(analyse *models.Analyse) error {
	args := m.Called(analyse)
	return args.Error(0)
}

func (m *MockAnalyseController) ExecuteAnalyse(externalId string) error {
	args := m.Called(externalId)
	return args.Error(0)
}

func (m *MockAnalyseController) RemoveAnalyse(toRemove *models.Analyse) error {
	args := m.Called(toRemove)
	return args.Error(0)
}

func (m *MockAnalyseController) PendingQueue() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
