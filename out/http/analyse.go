package http

import (
	"fmt"
	"io"
	"net/http"
	"pitzdev/web-service-gin/models"

	adapters "pitzdev/web-service-gin/out/adapters"
)

func GetScore(analyse *models.Analyse) (int, error) {
	fmt.Printf("[GetScore] Fetching score for Analyse %v\n", analyse.ID())

	url := "https://gingo.free.beeceptor.com/api/users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[GetScore] Error fetching score for Analyse %v: %v\n", analyse.ID(), err)
		return 0, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[GetScore] Error reading response body for Analyse %v: %v\n", analyse.ID(), err)
		return 0, err
	}

	return adapters.ParseScore(body)
}
