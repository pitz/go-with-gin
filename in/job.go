package in

import (
	"fmt"
	"pitzdev/web-service-gin/internal"
	"sync"
)

var (
	mu sync.Mutex
)

func ProcessQueue(c *internal.AnalyseController) {
	mu.Lock()
	defer mu.Unlock()

	for _, analyse := range c.PendingQueue() {
		err := c.ExecuteAnalyse(analyse)

		if err != nil {
			fmt.Println("[ProcessQueue] Error when processing job: ", err)
		}
	}
}
