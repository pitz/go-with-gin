package in

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pitzdev/web-service-gin/in"
	"pitzdev/web-service-gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseAnalyse(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Successful parsing", func(t *testing.T) {
		// given an valid payload
		postAnalyse := in.PostAnalyse{
			ExternalId: "external-id-1",
			Type:       "type-1",
			UserTaxId:  "tax-id-1",
		}
		jsonPayload, _ := json.Marshal(postAnalyse)
		req, _ := http.NewRequest(http.MethodPost, "/analyse", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		// and a valid request-context
		w := httptest.NewRecorder()
		requestContext, _ := gin.CreateTestContext(w)
		requestContext.Request = req

		// when calling the ParseAnalyse with the request-context
		analyse, err := in.ParseAnalyse(requestContext)

		// should return a valid parsed Analyse
		assert.Equal(t, postAnalyse.ExternalId, analyse.ExternalId)
		assert.Equal(t, postAnalyse.Type, analyse.Type)
		assert.Equal(t, postAnalyse.UserTaxId, analyse.UserTaxId)

		// and should not throw any error
		assert.NoError(t, err)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		// given an invalid payload
		jsonPayload := []byte(`{invalid-json}`)
		req, _ := http.NewRequest(http.MethodPost, "/analyse", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		// and a valid request-context
		w := httptest.NewRecorder()
		requestContext, _ := gin.CreateTestContext(w)
		requestContext.Request = req

		// when calling the ParseAnalyse with the request-context
		analyse, err := in.ParseAnalyse(requestContext)

		// should return nil
		assert.Error(t, err)
		assert.Equal(t, &models.Analyse{}, analyse)
	})
}
