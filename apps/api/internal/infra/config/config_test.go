package config

import (
	"os"
	"testing"
)

func TestLoad_Success(t *testing.T) {
	os.Setenv("APP_ENV", "local")
	os.Setenv("TELEGRAM_BOT_TOKEN", "test-token-123")
	os.Setenv("DATABASE_URL", "postgres://localhost:5432/test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("DATABASE_URL")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cfg.AppEnv != "local" {
		t.Errorf("expected AppEnv=local, got: %q", cfg.AppEnv)
	}
	if cfg.TelegramBotToken != "test-token-123" {
		t.Errorf("expected TelegramBotToken=test-token-123, got: %q", cfg.TelegramBotToken)
	}
	if cfg.DatabaseURL != "postgres://localhost:5432/test" {
		t.Errorf("expected DatabaseURL=postgres://localhost:5432/test, got: %q", cfg.DatabaseURL)
	}
}

func TestLoad_Defaults(t *testing.T) {
	os.Setenv("APP_ENV", "development")
	os.Setenv("TELEGRAM_BOT_TOKEN", "token")
	os.Setenv("DATABASE_URL", "postgres://db:5432/test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("DATABASE_URL")
	os.Unsetenv("PORT")
	os.Unsetenv("SCHEDULER_ENABLED")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cfg.Port != "8080" {
		t.Errorf("expected default Port=8080, got: %q", cfg.Port)
	}
	if cfg.SchedulerEnabled != "true" {
		t.Errorf("expected default SchedulerEnabled=true, got: %q", cfg.SchedulerEnabled)
	}
}

func TestLoad_MissingTelegramBotToken(t *testing.T) {
	os.Setenv("APP_ENV", "local")
	os.Setenv("DATABASE_URL", "postgres://db:5432/test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("DATABASE_URL")
	os.Unsetenv("TELEGRAM_BOT_TOKEN")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for missing TELEGRAM_BOT_TOKEN, got nil")
	}
}

func TestLoad_MissingDatabaseURL(t *testing.T) {
	os.Setenv("APP_ENV", "local")
	os.Setenv("TELEGRAM_BOT_TOKEN", "token")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	os.Unsetenv("DATABASE_URL")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for missing DATABASE_URL, got nil")
	}
}

func TestLoad_MissingAppEnv(t *testing.T) {
	os.Setenv("TELEGRAM_BOT_TOKEN", "token")
	os.Setenv("DATABASE_URL", "postgres://db:5432/test")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("DATABASE_URL")
	os.Unsetenv("APP_ENV")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for missing APP_ENV, got nil")
	}
}

func TestLoad_InvalidAppEnv(t *testing.T) {
	os.Setenv("APP_ENV", "staging")
	os.Setenv("TELEGRAM_BOT_TOKEN", "token")
	os.Setenv("DATABASE_URL", "postgres://db:5432/test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("DATABASE_URL")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid APP_ENV, got nil")
	}
}

func TestLoad_OptionalVarsNotRequired(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	os.Setenv("TELEGRAM_BOT_TOKEN", "token")
	os.Setenv("DATABASE_URL", "postgres://db:5432/test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("DATABASE_URL")
	os.Unsetenv("GEMINI_API_KEY")
	os.Unsetenv("OPENWEATHER_API_KEY")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error for missing optional vars, got: %v", err)
	}
	if cfg.GeminiAPIKey != "" {
		t.Errorf("expected empty GeminiAPIKey, got: %q", cfg.GeminiAPIKey)
	}
	if cfg.OpenWeatherAPIKey != "" {
		t.Errorf("expected empty OpenWeatherAPIKey, got: %q", cfg.OpenWeatherAPIKey)
	}
}
