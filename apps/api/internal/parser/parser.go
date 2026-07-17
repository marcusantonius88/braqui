package parser

import (
	"strings"
)

type Confidence int

const (
	ConfidenceLow    Confidence = 0
	ConfidenceMedium Confidence = 1
	ConfidenceHigh   Confidence = 2

	notMatched = "NOT_MATCHED"
)

func (c Confidence) String() string {
	switch c {
	case ConfidenceLow:
		return "low"
	case ConfidenceMedium:
		return "medium"
	case ConfidenceHigh:
		return "high"
	default:
		return "unknown"
	}
}

type ParseResult struct {
	Type       string
	Confidence Confidence
	Payload    map[string]any
}

type rule struct {
	eventType  string
	keywords   []string
	confidence Confidence
}

var rules = []rule{
	{eventType: "vomit", keywords: []string{"vomitou", "vomitando", "vômito", "vomitar", "vomitei"}, confidence: ConfidenceHigh},
	{eventType: "diarrhea", keywords: []string{"diarreia", "cocô mole", "cocô pastoso", "fezes moles"}, confidence: ConfidenceHigh},
	{eventType: "diarrhea", keywords: []string{"cocô", "fezes"}, confidence: ConfidenceLow},
	{eventType: "itching", keywords: []string{"coçando muito", "muita coceira"}, confidence: ConfidenceHigh},
	{eventType: "itching", keywords: []string{"coçando", "coceira", "se coçando"}, confidence: ConfidenceMedium},
	{eventType: "cough", keywords: []string{"tossindo", "tossiu", "tosse"}, confidence: ConfidenceHigh},
	{eventType: "fatigue", keywords: []string{"letárgico", "prostrado", "sem energia", "mole"}, confidence: ConfidenceHigh},
	{eventType: "fatigue", keywords: []string{"cansado", "cansaço"}, confidence: ConfidenceMedium},
	{eventType: "panting", keywords: []string{"respiração acelerada"}, confidence: ConfidenceHigh},
	{eventType: "panting", keywords: []string{"ofegante", "ofegando", "ofegar"}, confidence: ConfidenceMedium},
	{eventType: "medication_given", keywords: []string{"medicação", "remédio", "comprimido", "antipulga", "vermífugo", "simparic", "nexgard", "bravecto", "antibiótico", "anti-inflamatório"}, confidence: ConfidenceHigh},
	{eventType: "medication_given", keywords: []string{"tomou"}, confidence: ConfidenceMedium},
	{eventType: "weight_update", keywords: []string{"pesou", "pesagem", "pesar"}, confidence: ConfidenceHigh},
	{eventType: "weight_update", keywords: []string{"peso", "quilinhos", "quilos"}, confidence: ConfidenceMedium},
	{eventType: "vet_visit", keywords: []string{"veterinário", "veterinária", "vet"}, confidence: ConfidenceHigh},
	{eventType: "vet_visit", keywords: []string{"consulta", "consultar"}, confidence: ConfidenceMedium},
}

func Parse(text string) ParseResult {
	text = strings.TrimSpace(text)
	if text == "" {
		return ParseResult{Type: notMatched, Confidence: ConfidenceLow, Payload: map[string]any{}}
	}

	lower := strings.ToLower(text)

	var best ParseResult
	bestConfidence := Confidence(-1)

	for _, r := range rules {
		for _, kw := range r.keywords {
			if strings.Contains(lower, kw) {
				if r.confidence > bestConfidence {
					best = ParseResult{
						Type:       r.eventType,
						Confidence: r.confidence,
						Payload:    map[string]any{},
					}
					bestConfidence = r.confidence
				}
				break
			}
		}
	}

	if best.Type == "" {
		return ParseResult{Type: notMatched, Confidence: ConfidenceLow, Payload: map[string]any{}}
	}
	return best
}
