package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	domainai "github.com/marcusantonius88/braqui/apps/api/internal/domain/ai"
)

var geminiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

const defaultTimeout = 5 * time.Second

type appLogger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type GeminiProvider struct {
	apiKey string
	client *http.Client
	log    appLogger
}

func NewGeminiProvider(apiKey string, log appLogger) *GeminiProvider {
	return &GeminiProvider{
		apiKey: apiKey,
		client: &http.Client{Timeout: defaultTimeout},
		log:    log,
	}
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
}

type geminiCandidate struct {
	Content geminiContent `json:"content"`
}

const classificationPrompt = `Classifique a mensagem abaixo em um destes tipos: vomit, diarrhea, itching, cough, fatigue, panting, medication_given, weight_update, vet_visit.
Responda APENAS com um JSON válido no formato: {"type":"tipo","confidence":"high|medium|low","payload":{}}
Mensagem: `

func (p *GeminiProvider) Interpret(ctx context.Context, message string) (*domainai.InterpretationResult, error) {
	start := time.Now()
	p.log.Info("calling gemini", map[string]any{"message_length": len(message)})

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	prompt := classificationPrompt + message
	body := geminiRequest{
		Contents: []geminiContent{{
			Parts: []geminiPart{{Text: prompt}},
		}},
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		p.log.Error("gemini marshal error", map[string]any{"error": err.Error()})
		return notInterpreted(), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, geminiEndpoint+"?key="+p.apiKey, bytes.NewReader(reqBody))
	if err != nil {
		p.log.Error("gemini request error", map[string]any{"error": err.Error()})
		return notInterpreted(), nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		p.log.Error("gemini request failed", map[string]any{"error": err.Error(), "elapsed": time.Since(start).String()})
		return notInterpreted(), nil
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		p.log.Error("gemini read error", map[string]any{"error": err.Error()})
		return notInterpreted(), nil
	}

	if resp.StatusCode != http.StatusOK {
		p.log.Error("gemini api error", map[string]any{"status": resp.StatusCode, "body": string(raw), "elapsed": time.Since(start).String()})
		return notInterpreted(), nil
	}

	var gResp geminiResponse
	if err := json.Unmarshal(raw, &gResp); err != nil {
		p.log.Error("gemini unmarshal error", map[string]any{"error": err.Error()})
		return notInterpreted(), nil
	}

	if len(gResp.Candidates) == 0 || len(gResp.Candidates[0].Content.Parts) == 0 {
		p.log.Error("gemini empty response", nil)
		return notInterpreted(), nil
	}

	text := gResp.Candidates[0].Content.Parts[0].Text
	result, err := parseResult(text)
	if err != nil {
		p.log.Error("gemini invalid result", map[string]any{"error": err.Error(), "raw": text})
		return notInterpreted(), nil
	}

	p.log.Info("gemini success", map[string]any{"type": result.Type, "confidence": result.Confidence, "elapsed": time.Since(start).String()})
	return result, nil
}

func parseResult(text string) (*domainai.InterpretationResult, error) {
	text = cleanJSON(text)
	var result domainai.InterpretationResult
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		return nil, fmt.Errorf("parse result: %w", err)
	}
	if result.Payload == nil {
		result.Payload = map[string]any{}
	}
	if result.Type == "" {
		return nil, fmt.Errorf("empty type")
	}
	return &result, nil
}

func cleanJSON(text string) string {
	start := -1
	for i, c := range text {
		if c == '{' {
			start = i
			break
		}
	}
	if start == -1 {
		return "{}"
	}
	end := -1
	for i := len(text) - 1; i >= start; i-- {
		if text[i] == '}' {
			end = i
			break
		}
	}
	if end == -1 {
		return "{}"
	}
	return text[start : end+1]
}

func notInterpreted() *domainai.InterpretationResult {
	return &domainai.InterpretationResult{
		Type:       "NOT_INTERPRETED",
		Confidence: "low",
		Payload:    map[string]any{},
	}
}
