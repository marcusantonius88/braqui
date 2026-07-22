package climate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type openWeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

type OpenWeatherProvider struct {
	apiKey string
	client *http.Client
}

func NewOpenWeatherProvider(apiKey string) *OpenWeatherProvider {
	return &OpenWeatherProvider{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *OpenWeatherProvider) GetCurrentWeather(ctx context.Context, city string) (*WeatherData, error) {
	u := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		url.QueryEscape(city), p.apiKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openweather returned %d: %s", resp.StatusCode, string(body))
	}

	var owr openWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&owr); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &WeatherData{
		Temperature: owr.Main.Temp,
		Humidity:    owr.Main.Humidity,
		City:        owr.Name,
	}, nil
}
