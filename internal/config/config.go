package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Файл .env не найден", "error", err)
		panic(err)

	}

	config := &Config{
		DBHost:     viper.GetString("DBHost"),
		DBPort:     viper.GetString("DBPort"),
		DBUser:     viper.GetString("DBUser"),
		DBPassword: viper.GetString("DBPassword"),
		DBName:     viper.GetString("DBName"),
		//Port:       viper.GetString("Port"),
	}

	return config, nil

}
