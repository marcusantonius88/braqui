package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AppEnv            string
	TelegramBotToken  string
	DatabaseURL       string
	GeminiAPIKey      string
	OpenWeatherAPIKey string
	SchedulerEnabled  string
	Port              string
}

func Load() (*Config, error) {
	loadEnvFile(".env")

	cfg := &Config{
		AppEnv:            os.Getenv("APP_ENV"),
		TelegramBotToken:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		GeminiAPIKey:      os.Getenv("GEMINI_API_KEY"),
		OpenWeatherAPIKey: os.Getenv("OPENWEATHER_API_KEY"),
		SchedulerEnabled:  os.Getenv("SCHEDULER_ENABLED"),
		Port:              os.Getenv("PORT"),
	}

	setDefaults(cfg)

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func setDefaults(cfg *Config) {
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.SchedulerEnabled == "" {
		cfg.SchedulerEnabled = "true"
	}
}

func validate(cfg *Config) error {
	if cfg.TelegramBotToken == "" {
		return errors.New("missing required environment variable: TELEGRAM_BOT_TOKEN")
	}
	if cfg.DatabaseURL == "" {
		return errors.New("missing required environment variable: DATABASE_URL")
	}
	if cfg.AppEnv == "" {
		return errors.New("missing required environment variable: APP_ENV")
	}

	validEnvs := map[string]bool{
		"local":       true,
		"development": true,
		"production":  true,
	}
	if !validEnvs[cfg.AppEnv] {
		return fmt.Errorf("invalid APP_ENV: %q (must be local, development, or production)", cfg.AppEnv)
	}

	return nil
}

func loadEnvFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key != "" {
			os.Setenv(key, value)
		}
	}
}
