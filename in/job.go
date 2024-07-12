package in

import (
	"fmt"
	"pitzdev/web-service-gin/internal"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func ProcessQueue(analyseController *internal.AnalyseController) {
	fmt.Println("ProcessQueue: ", time.Now())
	fmt.Println("- Queue: ", len(*analyseController.AnalyseQueue()))

	mu.Lock()
	defer mu.Unlock()

	for _, analyse := range *analyseController.AnalyseQueue() {
		analyseController.ExecuteAnalyse(&analyse)
	}
}
