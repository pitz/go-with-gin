package main

import (
	"fmt"
	jobs "pitzdev/web-service-gin/in/jobs"
	"time"

	controllers "pitzdev/web-service-gin/controllers"
	httpIn "pitzdev/web-service-gin/in/http"
	httpOut "pitzdev/web-service-gin/out/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func routingOrchestrator(router *gin.Engine, httpServer *httpIn.AnalyseHandler) {
	router.POST("/analyse", httpServer.ExecuteAnalyse)

	// router.GET("/analyse", handlers.GetWorkout)
	// router.GET("/workouts/:id/executions", handlers.ListWorkoutExecutions)
	// router.GET("/workouts/:id/executions/:id", handlers.GetWorkoutExecution)
	// router.POST("/workouts/:id/executions/", handlers.GetWorkoutExecution)
}

func scheduleJobs(analyseController *controllers.AnalyseController) {
	c := cron.New()

	c.AddFunc("@every 10s", func() {
		fmt.Println("Cron job running at:", time.Now())
		jobs.ProcessQueue(analyseController)
	})

	c.Start()
}

func main() {
	httpClient := httpOut.New()
	analyseController := controllers.New(httpClient)
	httpServer := httpIn.New(analyseController)

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
