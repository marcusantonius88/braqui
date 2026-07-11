package migrations

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/braqui?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func TestCollectSQLFiles(t *testing.T) {
	files, err := CollectSQLFiles("testdata")
	if err != nil {
		t.Fatalf("collect sql files: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}
}

func TestRunMigrations(t *testing.T) {
	pool := getTestPool(t)
	ctx := context.Background()

	dir := "."
	err := Run(ctx, pool, dir)
	if err != nil {
		t.Fatalf("run migrations: %v", err)
	}

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM schema_migrations").Scan(&count)
	if err != nil {
		t.Fatalf("count migrations: %v", err)
	}
	if count < 1 {
		t.Fatal("expected at least 1 migration recorded")
	}
}

func TestRunMigrationsIsIdempotent(t *testing.T) {
	pool := getTestPool(t)
	ctx := context.Background()

	dir := "."
	if err := Run(ctx, pool, dir); err != nil {
		t.Fatalf("first run: %v", err)
	}

	if err := Run(ctx, pool, dir); err != nil {
		t.Fatalf("second run: %v", err)
	}
}

func TestReadFile(t *testing.T) {
	tmpDir := t.TempDir()
	sqlFile := filepath.Join(tmpDir, "test.sql")
	if err := os.WriteFile(sqlFile, []byte("SELECT 1"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	content, err := ReadFile(sqlFile)
	if err != nil {
		t.Fatalf("read file: %v", err)
	}
	if content != "SELECT 1" {
		t.Fatalf("expected 'SELECT 1', got '%s'", content)
	}
}
