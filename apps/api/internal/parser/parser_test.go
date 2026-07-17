package parser

import (
	"testing"
)

func TestParse_Empty(t *testing.T) {
	r := Parse("")
	if r.Type != notMatched {
		t.Fatalf("expected NOT_MATCHED, got %s", r.Type)
	}
}

func TestParse_VomitHigh(t *testing.T) {
	for _, msg := range []string{"vomitou", "Thor vomitou", "vomitando", "teve vômito", "vomitar", "vomitei"} {
		r := Parse(msg)
		if r.Type != "vomit" {
			t.Fatalf("msg=%q: expected vomit, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_DiarrheaHigh(t *testing.T) {
	for _, msg := range []string{"diarreia", "teve diarreia", "cocô mole", "cocô pastoso", "fezes moles"} {
		r := Parse(msg)
		if r.Type != "diarrhea" {
			t.Fatalf("msg=%q: expected diarrhea, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_DiarrheaLow(t *testing.T) {
	r := Parse("cocô")
	if r.Type != "diarrhea" {
		t.Fatalf("expected diarrhea, got %s", r.Type)
	}
	if r.Confidence != ConfidenceLow {
		t.Fatalf("expected low, got %s", r.Confidence)
	}
}

func TestParse_ItchingHigh(t *testing.T) {
	for _, msg := range []string{"coçando muito", "muita coceira"} {
		r := Parse(msg)
		if r.Type != "itching" {
			t.Fatalf("msg=%q: expected itching, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_ItchingMedium(t *testing.T) {
	for _, msg := range []string{"coçando", "coceira", "se coçando"} {
		r := Parse(msg)
		if r.Type != "itching" {
			t.Fatalf("msg=%q: expected itching, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceMedium {
			t.Fatalf("msg=%q: expected medium, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_CoughHigh(t *testing.T) {
	for _, msg := range []string{"tossindo", "tossiu", "tosse", "Thor tossiu"} {
		r := Parse(msg)
		if r.Type != "cough" {
			t.Fatalf("msg=%q: expected cough, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_FatigueHigh(t *testing.T) {
	for _, msg := range []string{"letárgico", "prostrado", "sem energia", "mole"} {
		r := Parse(msg)
		if r.Type != "fatigue" {
			t.Fatalf("msg=%q: expected fatigue, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_FatigueMedium(t *testing.T) {
	for _, msg := range []string{"cansado", "cansaço", "Thor cansado"} {
		r := Parse(msg)
		if r.Type != "fatigue" {
			t.Fatalf("msg=%q: expected fatigue, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceMedium {
			t.Fatalf("msg=%q: expected medium, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_PantingHigh(t *testing.T) {
	r := Parse("respiração acelerada")
	if r.Type != "panting" {
		t.Fatalf("expected panting, got %s", r.Type)
	}
	if r.Confidence != ConfidenceHigh {
		t.Fatalf("expected high, got %s", r.Confidence)
	}
}

func TestParse_PantingMedium(t *testing.T) {
	for _, msg := range []string{"ofegante", "ofegando", "ofegar"} {
		r := Parse(msg)
		if r.Type != "panting" {
			t.Fatalf("msg=%q: expected panting, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceMedium {
			t.Fatalf("msg=%q: expected medium, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_MedicationHigh(t *testing.T) {
	for _, msg := range []string{"medicação", "remédio", "comprimido", "antipulga", "vermífugo", "tomou simparic"} {
		r := Parse(msg)
		if r.Type != "medication_given" {
			t.Fatalf("msg=%q: expected medication_given, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_MedicationMedium(t *testing.T) {
	r := Parse("tomou")
	if r.Type != "medication_given" {
		t.Fatalf("expected medication_given, got %s", r.Type)
	}
	if r.Confidence != ConfidenceMedium {
		t.Fatalf("expected medium, got %s", r.Confidence)
	}
}

func TestParse_WeightHigh(t *testing.T) {
	for _, msg := range []string{"pesou", "pesagem", "pesar", "pesou 12kg"} {
		r := Parse(msg)
		if r.Type != "weight_update" {
			t.Fatalf("msg=%q: expected weight_update, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_WeightMedium(t *testing.T) {
	for _, msg := range []string{"peso", "quilinhos", "quilos", "ganhou peso"} {
		r := Parse(msg)
		if r.Type != "weight_update" {
			t.Fatalf("msg=%q: expected weight_update, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceMedium {
			t.Fatalf("msg=%q: expected medium, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_VetVisitHigh(t *testing.T) {
	for _, msg := range []string{"veterinário", "veterinária", "vet", "foi no veterinário"} {
		r := Parse(msg)
		if r.Type != "vet_visit" {
			t.Fatalf("msg=%q: expected vet_visit, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceHigh {
			t.Fatalf("msg=%q: expected high, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_VetVisitMedium(t *testing.T) {
	for _, msg := range []string{"consulta", "consultar", "consulta agendada"} {
		r := Parse(msg)
		if r.Type != "vet_visit" {
			t.Fatalf("msg=%q: expected vet_visit, got %s", msg, r.Type)
		}
		if r.Confidence != ConfidenceMedium {
			t.Fatalf("msg=%q: expected medium, got %s", msg, r.Confidence)
		}
	}
}

func TestParse_NotMatched(t *testing.T) {
	for _, msg := range []string{"bom dia", "qual a raça", "obrigado", "sim", "não", "Thor"} {
		r := Parse(msg)
		if r.Type != notMatched {
			t.Fatalf("msg=%q: expected NOT_MATCHED, got %s", msg, r.Type)
		}
	}
}

func TestParse_BestConfidenceWins(t *testing.T) {
	r := Parse("coçando muito")
	if r.Type != "itching" {
		t.Fatalf("expected itching, got %s", r.Type)
	}
	if r.Confidence != ConfidenceHigh {
		t.Fatalf("expected high, got %s", r.Confidence)
	}

	r2 := Parse("coçando")
	if r2.Type != "itching" {
		t.Fatalf("expected itching, got %s", r2.Type)
	}
	if r2.Confidence != ConfidenceMedium {
		t.Fatalf("expected medium, got %s", r2.Confidence)
	}
}

func TestConfidence_String(t *testing.T) {
	if ConfidenceLow.String() != "low" {
		t.Fatalf("expected low, got %s", ConfidenceLow.String())
	}
	if ConfidenceMedium.String() != "medium" {
		t.Fatalf("expected medium, got %s", ConfidenceMedium.String())
	}
	if ConfidenceHigh.String() != "high" {
		t.Fatalf("expected high, got %s", ConfidenceHigh.String())
	}
}
