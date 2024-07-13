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

func ProcessQueue(c *internal.AnalyseController) {
	mu.Lock()
	defer mu.Unlock()

	for _, analyse := range c.PendingQueue() {
		c.ExecuteAnalyse(analyse)
	}

	fmt.Println("DONE: ", time.Now())
}
