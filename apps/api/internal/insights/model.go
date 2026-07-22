package insights

import "time"

type Insight struct {
	Type        string
	Message     string
	GeneratedAt time.Time
}

const (
	InsightVomit       = "vomit_frequency"
	InsightItching     = "itching_frequency"
	InsightPanting     = "panting_frequency"
	InsightMedication  = "medication_absence"
)

const (
	EventVomit      = "vomit"
	EventItching    = "itching"
	EventPanting    = "panting"
	EventMedication = "medication_given"
)
