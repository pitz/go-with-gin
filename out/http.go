package out

import (
	"fmt"
	"io"
	"net/http"
	"pitzdev/web-service-gin/models"
	"time"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (h *Client) GetAdyenScore(analyse *models.Analyse, ch chan models.Score) {
	fmt.Printf("[GetAdyenScore] Fetching score for Analyse %v\n", analyse.ID())

	time.Sleep(1000 * time.Millisecond)
	url := "https://gingo.free.beeceptor.com/api/users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[GetAdyenScore] Error fetching score for Analyse %v: %v\n", analyse.ID(), err)
		ch <- models.Score{Error: err}
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[GetAdyenScore] Error reading response body for Analyse %v: %v\n", analyse.ID(), err)
		ch <- models.Score{Error: err}
		return
	}

	score, err := ParseScore(body)
	if err != nil {
		fmt.Printf("[GetAdyenScore] Error parsing response body for Analyse %v: %v\n It's going to use fallback.\n\n", analyse.ID(), err)
		ch <- models.Score{Error: nil, Score: 20}
		return
	}

	ch <- models.Score{Error: nil, Score: score}
}

func (h *Client) GetTransunionScore(analyse *models.Analyse, ch chan models.Score) {
	fmt.Printf("[GetTransunionScore] Fetching person background for Analyse %v\n", analyse.ID())

	time.Sleep(10000 * time.Millisecond)
	url := "https://gingo.free.beeceptor.com/api/users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[GetTransunionScore] Error fetching person background for Analyse %v: %v\n", analyse.ID(), err)
		ch <- models.Score{Error: err}
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[GetTransunionScore] Error reading response body for Analyse %v: %v\n", analyse.ID(), err)
		ch <- models.Score{Error: err}
		return
	}

	score, err := ParseScore(body)
	if err != nil {
		fmt.Printf("[GetTransunionScore] Error parsing response body for Analyse %v: %v\n It's going to use fallback.\n\n", analyse.ID(), err)
		ch <- models.Score{Error: nil, Score: 10}
		return
	}

	ch <- models.Score{Error: nil, Score: score}
}
