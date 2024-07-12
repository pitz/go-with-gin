package internal

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"pitzdev/web-service-gin/models"
	"pitzdev/web-service-gin/out"
)

type AnalyseControllerInterface interface {
	ScheduleExecution(analyse *models.Analyse) error
	ExecuteAnalyse(analyse *models.Analyse) error
	RemoveAnalyse(analyse *models.Analyse) error
	AnalyseQueue() *[]models.Analyse
}

type AnalyseController struct {
	httpClient   *out.Client
	analyseQueue []models.Analyse // I need to change this to be a MAP
}

func (c *AnalyseController) ScheduleExecution(analyse *models.Analyse) error {
	if err := c.validateAnalyse(analyse); err != nil {
		return err
	}

	analyse.SetID(uuid.New().String())
	c.analyseQueue = append(c.analyseQueue, *analyse)
	return nil
}

func (c *AnalyseController) validateAnalyse(analyse *models.Analyse) error {
	if analyse.ExternalId == "" {
		return errors.New("invalid ExternalID")
	}
	return nil
}

func (c *AnalyseController) ExecuteAnalyse(analyse *models.Analyse) error {
	fmt.Println("Execute: ", analyse.ID())

	score, err := c.httpClient.GetScore(analyse)
	if err != nil {
		return err
	}

	fmt.Println("Score: ", score)
	err = c.RemoveAnalyse(analyse)
	if err != nil {
		return err
	}

	return nil
}

func (c *AnalyseController) AnalyseQueue() *[]models.Analyse {
	return &c.analyseQueue
}

func (c *AnalyseController) RemoveAnalyse(analyseToRemove *models.Analyse) error {
	fmt.Println("Remove Analyse: ", analyseToRemove.ID())

	for i, analyse := range c.analyseQueue {
		if analyse.ID() == analyseToRemove.ID() {
			c.analyseQueue = append(c.analyseQueue[:i], c.analyseQueue[i+1:]...)
			return nil
		}
	}
	return errors.New("analyse not found")
}

func New(httpClient *out.Client) *AnalyseController {
	return &AnalyseController{
		httpClient:   httpClient,
		analyseQueue: []models.Analyse{},
	}
}
