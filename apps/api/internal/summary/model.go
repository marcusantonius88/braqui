package summary

import "fmt"

var typeLabels = map[string][2]string{
	"vomit":            {"episódio de vômito", "episódios de vômito"},
	"itching":          {"episódio de coceira", "episódios de coceira"},
	"panting":          {"episódio de ofegância", "episódios de ofegância"},
	"medication_given": {"medicação registrada", "medicações registradas"},
	"vet_visit":        {"consulta veterinária registrada", "consultas veterinárias registradas"},
}

func labelFor(typ string, count int) string {
	l, ok := typeLabels[typ]
	if !ok {
		return typ
	}
	if count == 1 {
		return l[0]
	}
	return l[1]
}

func pluralize(count int, singular, plural string) string {
	if count == 1 {
		return "1 " + singular
	}
	return itoa(count) + " " + plural
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}

type eventCount struct {
	Type  string
	Count int
}

var trackedTypes = []string{"vomit", "itching", "panting", "medication_given", "vet_visit"}

func formatSummary(petName string, counts map[string]int, total int) string {
	if total == 0 {
		return fmt.Sprintf("Nenhum evento foi registrado esta semana para o %s.", petName)
	}

	msg := fmt.Sprintf("📊 Resumo semanal do %s\n\n", petName)
	msg += "• " + pluralize(total, "evento registrado", "eventos registrados")

	var withEvents []eventCount
	for _, t := range trackedTypes {
		if c := counts[t]; c > 0 {
			withEvents = append(withEvents, eventCount{t, c})
		}
	}

	if len(withEvents) > 0 {
		msg += "\n\n"
		for i, ec := range withEvents {
			if i > 0 {
				msg += "\n"
			}
			msg += "• " + pluralize(ec.Count, labelFor(ec.Type, ec.Count), labelFor(ec.Type, ec.Count))
		}
	}

	msg += "\n\nContinue registrando eventos para que eu possa acompanhar melhor a saúde do " + petName + " 🐶"
	return msg
}
