package main

import (
	"pitzdev/web-service-gin/in"
	"pitzdev/web-service-gin/internal"
	"pitzdev/web-service-gin/out"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func routingOrchestrator(router *gin.Engine, http in.ServerInterface) {
	router.POST("/analyse", http.ExecuteAnalyse)
}

func scheduleJobs(controller *internal.AnalyseController) {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		in.ProcessQueue(controller)
	})

	c.Start()
}

func main() {
	adyenClient := out.NewAdyen()
	transunionClient := out.NewTransUnion()

	controller := internal.New(adyenClient, transunionClient)
	server := in.New(controller)

	scheduleJobs(controller)

	router := gin.Default()
	routingOrchestrator(router, server)

	router.Run("localhost:8080")
}
