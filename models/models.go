package models

type AnalyseType string

const (
	Lending    AnalyseType = "Lending"
	CreditCard AnalyseType = "CreditCard"
)

type Analyse struct {
	id         string
	ExternalId string
	UserTaxId  string
	Type       AnalyseType
}

func (analyse *Analyse) ID() string {
	return analyse.id
}

func (analyse *Analyse) SetID(analyseId string) {
	analyse.id = analyseId
}

type ScoreType string

const (
	TransUnion ScoreType = "TransUnion"
	Adyen      ScoreType = "Adyen"
)

type Score struct {
	Score int
	Error error
	Type  ScoreType
}
