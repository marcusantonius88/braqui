package climate

import "context"

type WeatherData struct {
	Temperature float64
	Humidity    int
	City        string
}

type Provider interface {
	GetCurrentWeather(ctx context.Context, city string) (*WeatherData, error)
}
