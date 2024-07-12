package main

import (
	"fmt"

	"pitzdev/web-service-gin/in"
	"pitzdev/web-service-gin/internal"
	"pitzdev/web-service-gin/out"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func routingOrchestrator(router *gin.Engine, httpServer *in.Http) {
	router.POST("/analyse", httpServer.ExecuteAnalyse)
}

func scheduleJobs(analyseController *internal.AnalyseController) {
	c := cron.New()

	c.AddFunc("@every 10s", func() {
		fmt.Println("Cron job running at:", time.Now())
		in.ProcessQueue(analyseController)
	})

	c.Start()
}

func main() {
	httpClient := out.New()
	analyseController := internal.New(httpClient)
	httpServer := in.New(analyseController)

	testingGoroutine()
	scheduleJobs(analyseController)

	router := gin.Default()
	routingOrchestrator(router, httpServer)

	router.Run("localhost:8080")
}

// Lending Orchestrator
// - Receives the tax-id (document)
// - Break the flow into steps (read it from a file)
// - Get the flow to execute (get it from a map - try to use a random based on experiment)
// - Execute step
// - Send to the next step
