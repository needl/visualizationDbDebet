package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Port       string
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetDefault("DB_SSLMODE", "disable")

	if _, err := os.Stat(".env"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Warn("config file .env not found, using environment variables")
		} else {
			return nil, fmt.Errorf("check config file .env: %w", err)
		}
	} else if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfg := &Config{
		DBHost:     strings.TrimSpace(v.GetString("DB_HOST")),
		DBPort:     strings.TrimSpace(v.GetString("DB_PORT")),
		DBUser:     strings.TrimSpace(v.GetString("DB_USER")),
		DBPassword: v.GetString("DB_PASSWORD"),
		DBName:     strings.TrimSpace(v.GetString("DB_NAME")),
		DBSSLMode:  strings.TrimSpace(v.GetString("DB_SSLMODE")),
		Port:       strings.TrimSpace(v.GetString("PORT")),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	required := map[string]string{
		"DB_HOST": c.DBHost,
		"DB_PORT": c.DBPort,
		"DB_USER": c.DBUser,
		"DB_NAME": c.DBName,
		"PORT":    c.Port,
	}

	var missing []string
	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required config values: %s", strings.Join(missing, ", "))
	}

	if strings.TrimSpace(c.DBSSLMode) == "" {
		c.DBSSLMode = "disable"
	}

	return nil
}

func (c *Config) DBConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}
