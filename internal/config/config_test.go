package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_FromEnvWithoutFile(t *testing.T) {
	dir := t.TempDir()
	withWorkingDir(t, dir)
	setBaseEnv(t)

	t.Setenv("DB_HOST", " localhost ")
	t.Setenv("DB_PORT", " 5432 ")
	t.Setenv("DB_USER", " postgres ")
	t.Setenv("DB_PASSWORD", "secret")
	t.Setenv("DB_NAME", " app ")
	t.Setenv("DB_SSLMODE", "")
	t.Setenv("PORT", " 8080 ")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.DBHost != "localhost" {
		t.Fatalf("unexpected DBHost: %q", cfg.DBHost)
	}
	if cfg.DBPort != "5432" {
		t.Fatalf("unexpected DBPort: %q", cfg.DBPort)
	}
	if cfg.DBUser != "postgres" {
		t.Fatalf("unexpected DBUser: %q", cfg.DBUser)
	}
	if cfg.DBName != "app" {
		t.Fatalf("unexpected DBName: %q", cfg.DBName)
	}
	if cfg.DBSSLMode != "disable" {
		t.Fatalf("expected DBSSLMode=disable, got %q", cfg.DBSSLMode)
	}
	if cfg.Port != "8080" {
		t.Fatalf("unexpected Port: %q", cfg.Port)
	}
}

func TestLoad_FromEnvFile(t *testing.T) {
	dir := t.TempDir()
	withWorkingDir(t, dir)
	setBaseEnv(t)

	envPath := filepath.Join(dir, ".env")
	data := strings.Join([]string{
		"DB_HOST=127.0.0.1",
		"DB_PORT=5432",
		"DB_USER=appuser",
		"DB_PASSWORD=secret",
		"DB_NAME=appdb",
		"PORT=9090",
		"",
	}, "\n")

	if err := os.WriteFile(envPath, []byte(data), 0o600); err != nil {
		t.Fatalf("failed to create .env: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.DBHost != "127.0.0.1" || cfg.DBPort != "5432" || cfg.DBUser != "appuser" || cfg.DBName != "appdb" || cfg.Port != "9090" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
	if cfg.DBSSLMode != "disable" {
		t.Fatalf("expected DBSSLMode=disable, got %q", cfg.DBSSLMode)
	}
}

func TestLoad_MissingRequiredValues(t *testing.T) {
	dir := t.TempDir()
	withWorkingDir(t, dir)
	setBaseEnv(t)

	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	// DB_USER missing
	t.Setenv("DB_NAME", "app")
	t.Setenv("PORT", "8080")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "missing required config values") {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(err.Error(), "DB_USER") {
		t.Fatalf("expected DB_USER in error, got %v", err)
	}
}

func TestDBConnString(t *testing.T) {
	cfg := &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "secret",
		DBName:     "app",
		DBSSLMode:  "disable",
	}

	got := cfg.DBConnString()
	want := "host=localhost port=5432 user=postgres password=secret dbname=app sslmode=disable"
	if got != want {
		t.Fatalf("unexpected conn string: %q", got)
	}
}

func withWorkingDir(t *testing.T, dir string) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
}

func setBaseEnv(t *testing.T) {
	t.Helper()
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE",
		"PORT",
	}
	for _, k := range keys {
		t.Setenv(k, "")
	}
}
