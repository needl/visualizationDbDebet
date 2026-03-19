package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"io"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/config"
	"visualizationBdDebet/internal/debet"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(err.Error(), "info", "cfg, err := config.LoadConfig() str 16")
	}

	db, err := sqlx.Connect("postgres", cfg.DBHost)
	if err != nil {
		slog.Error(err.Error(), "info", "db, err := sqlx.Connect(postgres, cfg.DBHost) str 21")
	}
	defer db.Close()
	debetRepo := debet.NewRepository(db)

	debetService := debet.NewService(debetRepo)
}
