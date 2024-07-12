package in

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pitzdev/web-service-gin/in"
	"pitzdev/web-service-gin/models"
	"pitzdev/web-service-gin/tests/fixtures"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExecuteAnalyse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Testing valid analyse", func(t *testing.T) {
		// setup
		mockController := new(fixtures.MockAnalyseController)
		handler := in.New(mockController)

		// given a valid analyse
		postAnalyse := models.Analyse{
			ExternalId: "external-id-1",
			UserTaxId:  "tax-id-1",
			Type:       "type-1",
		}
		jsonPayload, _ := json.Marshal(postAnalyse)
		req, _ := http.NewRequest(http.MethodPost, "/analyse", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(w)
		context.Request = req
		mockController.On("ScheduleExecution", mock.AnythingOfType("*models.Analyse")).Return(nil)

		// when calling ExecuteAnalyse
		handler.ExecuteAnalyse(context)

		// should return a valid response
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Analyse scheduled with success!")

		// and it should schedule the execution of the analysis
		mockController.AssertCalled(t, "ScheduleExecution", mock.AnythingOfType("*models.Analyse"))
	})

	t.Run("Testing an invalid analyse", func(t *testing.T) {
		// setup
		mockController := new(fixtures.MockAnalyseController)
		handler := in.New(mockController)

		// given an invalid analyse
		jsonPayload := []byte(`{banana: 1}`)
		req, _ := http.NewRequest(http.MethodPost, "/analyse", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(w)
		context.Request = req

		// when calling ExecuteAnalyse
		handler.ExecuteAnalyse(context)

		// should return error
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("Testing an error when scheduling the analyse", func(t *testing.T) {
		// setup
		mockController := new(fixtures.MockAnalyseController)
		mockController.On("ScheduleExecution", mock.AnythingOfType("*models.Analyse")).Return(errors.New("scheduling error"))
		handler := in.New(mockController)

		// given an valid analyse
		postAnalyse := models.Analyse{
			ExternalId: "external-id-1",
			UserTaxId:  "tax-id-1",
			Type:       "type-1",
		}
		jsonPayload, _ := json.Marshal(postAnalyse)
		req, _ := http.NewRequest(http.MethodPost, "/analyse", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(w)
		context.Request = req

		// when calling ExecuteAnalyse
		handler.ExecuteAnalyse(context)

		// should return error because was not able to schedule the analyse
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
		mockController.AssertCalled(t, "ScheduleExecution", mock.AnythingOfType("*models.Analyse"))
	})
}
