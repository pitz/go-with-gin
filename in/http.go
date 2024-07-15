package in

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pitzdev/web-service-gin/internal"
)

type ServerInterface interface {
	ExecuteAnalyse(context *gin.Context)
}

type Server struct {
	controller internal.AnalyseControllerInterface
}

func (h *Server) ExecuteAnalyse(context *gin.Context) {
	analyse, parsingErr := ParseAnalyse(context)
	if parsingErr != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": parsingErr.Error()})
		return
	}

	if scheduleError := h.controller.ScheduleExecution(analyse); scheduleError != nil {
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

func New(c internal.AnalyseControllerInterface) *Server {
	return &Server{controller: c}
}
