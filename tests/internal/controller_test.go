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
	Type:       "CreditCard",
}

var anotherValidAnalyse = &models.Analyse{
	ExternalId: "external-id-2",
	UserTaxId:  "tax-id-2",
	Type:       "Lending",
}

func TestAnalyseController_ScheduleExecution(t *testing.T) {
	mockedAdyenClient := new(fixtures.MockClient)
	mockedTransUnionClient := new(fixtures.MockClient)

	t.Run("Scheduling a valid Analyse", func(t *testing.T) {
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)

		// when scheduling two valid analyses
		err1 := controller.ScheduleExecution(validAnalyse)
		err2 := controller.ScheduleExecution(anotherValidAnalyse)

		// should return no errors
		assert.NoError(t, err1)
		assert.NoError(t, err2)

		// and the analyses should be present on PendingQueue()
		assert.Contains(t, controller.PendingQueue(), validAnalyse.ExternalId)
		assert.Contains(t, controller.PendingQueue(), anotherValidAnalyse.ExternalId)
	})

	t.Run("Re-scheduling a valid Analyse", func(t *testing.T) {
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)

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
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)

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

func TestAnalyseController_RemoveAnalyse(t *testing.T) {
	mockedAdyenClient := new(fixtures.MockClient)
	mockedTransUnionClient := new(fixtures.MockClient)

	t.Run("Removing a valid Analyse", func(t *testing.T) {
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)
		controller.ScheduleExecution(validAnalyse)

		// when removing a valid analyse
		err := controller.RemoveAnalyse(validAnalyse)

		// should return no errors
		assert.NoError(t, err)

		// and the analyse should no longer be present on PendingQueue()
		assert.NotContains(t, controller.PendingQueue(), validAnalyse.ExternalId)
	})

	t.Run("Removing an already removed Analyse", func(t *testing.T) {
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)
		controller.ScheduleExecution(validAnalyse)
		controller.RemoveAnalyse(validAnalyse)

		// when trying to remove the same analyse again
		err := controller.RemoveAnalyse(validAnalyse)

		// should return an error
		assert.Error(t, err, "analyse not found")
	})

	t.Run("Removing a non-existent Analyse", func(t *testing.T) {
		// setup
		controller := internal.New(mockedAdyenClient, mockedTransUnionClient)

		// when trying to remove the non-existent analyse
		err := controller.RemoveAnalyse(anotherValidAnalyse)

		// should return an error
		assert.Error(t, err, "analyse not found")
	})
}
