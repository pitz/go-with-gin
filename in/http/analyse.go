package http

import (
	"net/http"
	controllers "pitzdev/web-service-gin/controllers"
	adapters "pitzdev/web-service-gin/in/adapters"

	"github.com/gin-gonic/gin"
)

func ExecuteAnalyse(context *gin.Context) {
	analyse, parsingErr := adapters.ParseAnalyse(context)
	if parsingErr != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": parsingErr.Error()})
		return
	}

	scheduleError := controllers.ScheduleExecution(analyse)
	if scheduleError != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": scheduleError.Error()})
		return
	}

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "Analyse scheduled with success!"},
	)
}
