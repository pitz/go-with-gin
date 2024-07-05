package controllers

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	models "pitzdev/web-service-gin/models"
	http "pitzdev/web-service-gin/out/http"
)

var AnalyseQueue = []models.Analyse{}

func validateAnalyse(analyse *models.Analyse) error {
	if analyse.ExternalId == "" {
		return errors.New("invalid ExternalID")
	}

	return nil
}

func ScheduleExecution(analyse *models.Analyse) error {
	err := validateAnalyse(analyse)
	if err != nil {
		return err
	}

	analyse.SetID(uuid.New().String())
	AnalyseQueue = append(AnalyseQueue, *analyse)

	return nil
}

func ExecuteAnalyse(analyse *models.Analyse) error {
	fmt.Println("Execute: ", analyse.ID())

	score, err := http.GetScore(analyse)
	if err != nil {
		return err
	}

	// tmp
	fmt.Println("Score: ", score)

	return nil
}
