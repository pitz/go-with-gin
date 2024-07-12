package in

import (
	"github.com/gin-gonic/gin"
	"pitzdev/web-service-gin/models"
)

func ParseAnalyse(context *gin.Context) (*models.Analyse, error) {
	var postAnalyse PostAnalyse

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
