package climate

type RiskLevel int

const (
	RiskNone     RiskLevel = 0
	RiskHigh     RiskLevel = 1
	RiskCritical RiskLevel = 2
)

func (r RiskLevel) String() string {
	switch r {
	case RiskHigh:
		return "alta"
	case RiskCritical:
		return "crítica"
	default:
		return "nenhum"
	}
}

func EvaluateRisk(temp float64) RiskLevel {
	if temp >= 35 {
		return RiskCritical
	}
	if temp >= 30 {
		return RiskHigh
	}
	return RiskNone
}
