package climate

import "fmt"

func FormatAlert(city string, temp float64, risk RiskLevel) string {
	msg := fmt.Sprintf("🐶 Atenção\n\nHoje está muito quente em %s (%.0f°C).\n\nCães braquicefálicos podem sofrer mais nesses dias.", city, temp)

	if risk == RiskCritical {
		msg += "\n\n⚠️ Calor crítico! Evite passeios e mantenha seu pet em local arejado."
	}

	return msg
}
