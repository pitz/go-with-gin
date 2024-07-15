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
	ExecuteAnalyse(externalId string) error
	PendingQueue() []string
	RemoveAnalyse(toRemove *models.Analyse) error
}

type AnalyseController struct {
	adyenClient      out.ClientInterface
	transUnionClient out.ClientInterface
	analyseQueue     map[string]models.Analyse
}

func (c *AnalyseController) ScheduleExecution(analyse *models.Analyse) error {
	if err := c.validateAnalyse(analyse); err != nil {
		return err
	}

	analyse.SetID(uuid.New().String())
	c.analyseQueue[analyse.ExternalId] = *analyse

	return nil
}

func (c *AnalyseController) validateAnalyse(analyse *models.Analyse) error {
	if analyse.ExternalId == "" {
		return errors.New("invalid ExternalID")
	}

	if _, present := c.analyseQueue[analyse.ExternalId]; present {
		return errors.New("duplicated ExternalID")
	}

	return nil
}

func (c *AnalyseController) ExecuteAnalyse(externalId string) error {
	fmt.Println("Execute: ", externalId)

	analyse, present := c.analyseQueue[externalId]
	if !present {
		return errors.New("analyse not found")
	}

	ch := make(chan models.Score)
	go c.adyenClient.GetScore(&analyse, ch)
	go c.transUnionClient.GetScore(&analyse, ch)

	score := make([]models.Score, 2)
	score[0] = <-ch
	score[1] = <-ch
	fmt.Println("Scores: ", score)

	err := c.RemoveAnalyse(&analyse)
	if err != nil {
		return err
	}

	return nil
}

func (c *AnalyseController) PendingQueue() []string {
	keys := make([]string, len(c.analyseQueue))

	i := 0
	for k := range c.analyseQueue {
		keys[i] = k
		i++
	}

	return keys
}

func (c *AnalyseController) RemoveAnalyse(toRemove *models.Analyse) error {
	fmt.Println("Remove Analyse: ", toRemove.ID())

	if _, ok := c.analyseQueue[toRemove.ExternalId]; ok {
		delete(c.analyseQueue, toRemove.ExternalId)
		return nil
	}

	return errors.New("analyse not found")
}

func New(
	adyenClient out.ClientInterface,
	transUnionClient out.ClientInterface,
) *AnalyseController {
	return &AnalyseController{
		adyenClient:      adyenClient,
		transUnionClient: transUnionClient,
		analyseQueue:     make(map[string]models.Analyse),
	}
}
