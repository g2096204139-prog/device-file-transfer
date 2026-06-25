package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerHost      string
	ServerPort      string
	UploadDir       string
	AccessToken     string
	AuthEnabled     bool
	MaxUploadSizeMB int64
	AllowDelete     bool
	LogDir          string
}

func Load() Config {
	loadDotEnv(".env")

	return Config{
		ServerHost:      getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort:      getEnv("SERVER_PORT", "8000"),
		UploadDir:       getEnv("UPLOAD_DIR", "uploads"),
		AccessToken:     getEnv("ACCESS_TOKEN", "change-this-token"),
		AuthEnabled:     getEnvBool("AUTH_ENABLED", true),
		MaxUploadSizeMB: getEnvInt64("MAX_UPLOAD_SIZE_MB", 500),
		AllowDelete:     getEnvBool("ALLOW_DELETE", true),
		LogDir:          getEnv("LOG_DIR", "logs"),
	}
}

func (c Config) MaxUploadBytes() int64 {
	if c.MaxUploadSizeMB <= 0 {
		return 0
	}
	return c.MaxUploadSizeMB * 1024 * 1024
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)

		if key == "" {
			continue
		}

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		_ = os.Setenv(key, value)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvInt64(key string, fallback int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}
	return parsed
}
