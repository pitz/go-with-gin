package out

import (
	"encoding/json"
	"math/rand"
)

func ParseScore(body []byte) (int, error) {
	var scoreResp Score
	err := json.Unmarshal(body, &scoreResp)
	if err != nil {
		return 0, err
	}

	score := scoreResp.Score
	if score != 0 {
		return score, nil
	}

	randomNumber := rand.Intn(10)
	return randomNumber, nil
}
