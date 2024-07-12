package in

import (
	"pitzdev/web-service-gin/models"
)

type PostAnalyse struct {
	ExternalId string             `json:"externalId"`
	UserTaxId  string             `json:"userTaxId"`
	Type       models.AnalyseType `json:"type"`
}
