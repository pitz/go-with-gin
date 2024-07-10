package http

import (
	"net/http"
	"pitzdev/web-service-gin/controllers"
	"pitzdev/web-service-gin/in/adapters"

	"github.com/gin-gonic/gin"
)

type AnalyseHandler struct {
	controller *controllers.AnalyseController
}

func (h *AnalyseHandler) ExecuteAnalyse(context *gin.Context) {
	analyse, parsingErr := adapters.ParseAnalyse(context)
	if parsingErr != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": parsingErr.Error()})
		return
	}

	scheduleError := h.controller.ScheduleExecution(analyse)
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

func New(controller *controllers.AnalyseController) *AnalyseHandler {
	return &AnalyseHandler{controller: controller}
}
