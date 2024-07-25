package out

import (
	"fmt"
	"io"
	"net/http"
	"pitzdev/web-service-gin/models"
	"time"
)

type ClientInterface interface {
	GetScore(analyse *models.Analyse, ch chan<- models.Score)
}

type AdyenClient struct{}

func NewAdyen() *AdyenClient {
	return &AdyenClient{}
}

func (h *AdyenClient) GetScore(analyse *models.Analyse, ch chan<- models.Score) {
	time.Sleep(1000 * time.Millisecond) // temp

	score := models.Score{Type: models.Adyen}
	url := "https://gingo.free.beeceptor.com/api/users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[AdyenClient] Error fetching score for Analyse %v: %v\n", analyse.ID(), err)
		score.Error = err
		ch <- score
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[AdyenClient] Error reading response body for Analyse %v: %v\n", analyse.ID(), err)
		score.Error = err
		ch <- score
		return
	}

	adyenScore, err := ParseScore(body)
	if err != nil {
		fmt.Printf("[AdyenClient] Error parsing response body for Analyse %v: %v\n It's going to use fallback.\n\n", analyse.ID(), err)
		score.Score = 20
		ch <- score
		return
	}

	score.Score = adyenScore
	ch <- score
}

type TransUnionClient struct{}

func NewTransUnion() *TransUnionClient {
	return &TransUnionClient{}
}

func (h *TransUnionClient) GetScore(analyse *models.Analyse, ch chan<- models.Score) {
	time.Sleep(5000 * time.Millisecond)

	score := models.Score{Type: models.TransUnion}
	url := "https://gingo.free.beeceptor.com/api/users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[TransUnionClient] Error fetching person background for Analyse %v: %v\n", analyse.ID(), err)
		score.Error = err
		ch <- score
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[TransUnionClient] Error reading response body for Analyse %v: %v\n", analyse.ID(), err)
		score.Error = err
		ch <- score
		return
	}

	transUnionScore, err := ParseScore(body)
	if err != nil {
		fmt.Printf("[TransUnionClient] Error parsing response body for Analyse %v: %v\n It's going to use fallback.\n\n", analyse.ID(), err)
		score.Score = 10
		ch <- score
		return
	}

	ch <- models.Score{Error: nil, Score: transUnionScore}
}

// This should be in diff files.
