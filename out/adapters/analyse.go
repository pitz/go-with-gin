package adapters

import (
	"encoding/json"
	"fmt"
	"math/rand"

	schemas "pitzdev/web-service-gin/out/schemas"
)

func ParseScore(body []byte) (int, error) {
	var scoreResp schemas.Score
	err := json.Unmarshal(body, &scoreResp)
	if err != nil {
		fmt.Printf("[ParseScore] Error decoding JSON for Analyse: %v\n", err)
		return 0, err
	}

	score := scoreResp.Score
	if score != 0 {
		return score, nil
	}

	randomNumber := rand.Intn(10)
	return randomNumber, nil
}
