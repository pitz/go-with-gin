package httpClient

import (
	"fmt"
	"io"
	"net/http"
	"pitzdev/web-service-gin/models"

	adapters "pitzdev/web-service-gin/out/adapters"
)

type HttpClient struct{}

func New() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) GetScore(analyse *models.Analyse) (int, error) {
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

	score, err := adapters.ParseScore(body)
	if err != nil {
		fmt.Printf("[GetScore] Error parsing response body for Analyse %v: %v\n. It's going to use fallback.", analyse.ID(), err)
		return 0, nil
	}

	return score, nil
}
