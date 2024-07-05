package main

import (
	"fmt"
	http "pitzdev/web-service-gin/in/http"
	jobs "pitzdev/web-service-gin/in/jobs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func routingOrchestrator(router *gin.Engine) {
	router.POST("/analyse", http.ExecuteAnalyse)

	// router.GET("/analyse", handlers.GetWorkout)
	// router.GET("/workouts/:id/executions", handlers.ListWorkoutExecutions)
	// router.GET("/workouts/:id/executions/:id", handlers.GetWorkoutExecution)
	// router.POST("/workouts/:id/executions/", handlers.GetWorkoutExecution)
}

func scheduleJobs() {
	c := cron.New()

	c.AddFunc("@every 10s", func() {
		fmt.Println("Cron job running at:", time.Now())
		jobs.ProcessQueue()
	})

	c.Start()
}

func main() {
	testingGoroutine()
	scheduleJobs()

	router := gin.Default()
	routingOrchestrator(router)

	router.Run("localhost:8080")
}

// Lending Orchestrator
// - Receives the tax-id (document)
// - Break the flow into steps (read it from a file)
// - Get the flow to execute (get it from a map - try to use a random based on experiment)
// - Execute step
// - Send to the next step
