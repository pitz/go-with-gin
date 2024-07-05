package adapters

import (
	schemas "pitzdev/web-service-gin/in/schemas"
	models "pitzdev/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func ParseAnalyse(context *gin.Context) (*models.Analyse, error) {
	var postAnalyse schemas.PostAnalyse

	err := context.BindJSON(&postAnalyse)
	if err != nil {
		return &models.Analyse{}, err
	}

	var analyse models.Analyse
	analyse.ExternalId = postAnalyse.ExternalId
	analyse.Type = postAnalyse.Type
	analyse.UserTaxId = postAnalyse.UserTaxId
	return &analyse, nil
}
