package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	t.Setenv("SERVER_PORT", "")
	cfg := Load()

	if cfg.ServerPort != "8000" {
		t.Fatalf("expected default port 8000, got %s", cfg.ServerPort)
	}

	if cfg.MaxUploadSizeMB != 500 {
		t.Fatalf("expected default max upload 500, got %d", cfg.MaxUploadSizeMB)
	}
}

func TestMaxUploadBytesUnlimited(t *testing.T) {
	cfg := Config{MaxUploadSizeMB: 0}
	if cfg.MaxUploadBytes() != 0 {
		t.Fatal("expected 0 bytes to mean unlimited")
	}
}

func TestEnvConfig(t *testing.T) {
	t.Setenv("SERVER_PORT", "9000")
	t.Setenv("MAX_UPLOAD_SIZE_MB", "123")
	t.Setenv("AUTH_ENABLED", "false")

	cfg := Load()

	if cfg.ServerPort != "9000" {
		t.Fatalf("expected port 9000, got %s", cfg.ServerPort)
	}

	if cfg.MaxUploadSizeMB != 123 {
		t.Fatalf("expected max upload 123, got %d", cfg.MaxUploadSizeMB)
	}

	if cfg.AuthEnabled {
		t.Fatal("expected auth disabled")
	}

	_ = os.Getenv("SERVER_PORT")
}

func TestLoadDotEnvFile(t *testing.T) {
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Chdir(oldWd)
	}()

	envPath := filepath.Join(tempDir, ".env")
	if err := os.WriteFile(envPath, []byte("SERVER_PORT=9100\nACCESS_TOKEN=local-token\n"), 0600); err != nil {
		t.Fatal(err)
	}

	_ = os.Unsetenv("SERVER_PORT")
	_ = os.Unsetenv("ACCESS_TOKEN")
	defer func() {
		_ = os.Unsetenv("SERVER_PORT")
		_ = os.Unsetenv("ACCESS_TOKEN")
	}()

	cfg := Load()
	if cfg.ServerPort != "9100" {
		t.Fatalf("expected .env port 9100, got %s", cfg.ServerPort)
	}
	if cfg.AccessToken != "local-token" {
		t.Fatalf("expected .env access token, got %s", cfg.AccessToken)
	}
}
