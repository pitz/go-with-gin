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

func ProcessQueue(analyseController *controllers.AnalyseController) {
	fmt.Println("ProcessQueue: ", time.Now())
	fmt.Println("- Queue: ", len(*analyseController.AnalyseQueue()))

	mu.Lock()
	defer mu.Unlock()

	for _, analyse := range *analyseController.AnalyseQueue() {
		analyseController.ExecuteAnalyse(&analyse)
	}

	// controllers.AnalyseQueue = []models.Analyse{}
}
