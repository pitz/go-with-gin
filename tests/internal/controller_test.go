package internal

import (
	"pitzdev/web-service-gin/internal"
	"pitzdev/web-service-gin/models"
	"pitzdev/web-service-gin/tests/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAnalyse = &models.Analyse{
	ExternalId: "external-id-1",
	UserTaxId:  "tax-id-1",
	Type:       "type-1",
}

func TestAnalyseController_ScheduleExecution(t *testing.T) {
	mockedAdyenClient := new(fixtures.MockClient)
	mockedTransUnionClient := new(fixtures.MockClient)
	controller := internal.New(mockedAdyenClient, mockedTransUnionClient)

	t.Run("Scheduling a valid Analyse", func(t *testing.T) {
		// when schedulling a valid analyse
		err := controller.ScheduleExecution(validAnalyse)

		// should return no errors
		assert.NoError(t, err)

		// and the analyse should be present on PendingQueue()
		assert.Contains(t, controller.PendingQueue(), validAnalyse.ExternalId)
	})

	t.Run("Re-scheduling a valid Analyse", func(t *testing.T) {
		// when scheduling a valid analyse
		err1 := controller.ScheduleExecution(validAnalyse)

		// should return no errors for the first call
		assert.NoError(t, err1)

		// and scheduling the same analyse again
		err2 := controller.ScheduleExecution(validAnalyse)

		// should return an error for the second call
		assert.Error(t, err2, "duplicated ExternalID")
	})

	t.Run("Scheduling an invalid Analyse", func(t *testing.T) {
		// given a analyse without external-id
		var invalid = &models.Analyse{
			ExternalId: "",
			UserTaxId:  "tax-id-1",
			Type:       "type-1",
		}

		// when scheduling twice the same analyse
		err := controller.ScheduleExecution(invalid)

		// should return no errors for the first call
		assert.Error(t, err, "invalid ExternalID")
	})
}
