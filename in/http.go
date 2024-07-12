package in

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pitzdev/web-service-gin/internal"
)

type Http struct {
	controller internal.AnalyseControllerInterface
}

func (h *Http) ExecuteAnalyse(context *gin.Context) {
	analyse, parsingErr := ParseAnalyse(context)
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

func New(controller internal.AnalyseControllerInterface) *Http {
	return &Http{controller: controller}
}
