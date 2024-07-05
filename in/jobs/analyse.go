package jobs

import (
	"fmt"
	"sync"
	"time"

	controllers "pitzdev/web-service-gin/controllers"
)

var (
	mu sync.Mutex
)

func ProcessQueue() {
	fmt.Println("ProcessQueue: ", time.Now())

	mu.Lock()
	defer mu.Unlock()

	for _, analyse := range controllers.AnalyseQueue {
		controllers.ExecuteAnalyse(&analyse)
	}

	// controllers.AnalyseQueue = []models.Analyse{}
}
