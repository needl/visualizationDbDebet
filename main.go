package main

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"visualizationBdDebet/internal/config"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/delivery"
	"visualizationBdDebet/internal/delivery/handler"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		panic(err)
	}

	db, err := sqlx.Connect("postgres", cfg.DBConnString())
	if err != nil {
		slog.Error("Failed to connect to db", "error", err)
		panic(err)
	}
	defer db.Close()

	debetRepo := debet.NewRepository(db)
	debetService := debet.NewService(debetRepo)

	debetHandler := handler.NewDebetHandler(debetService)

	r := delivery.NewRouter(debetHandler)

	srv := http.Server{
		Handler:      r,
		Addr:         ":" + cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start server", "error", err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown Failed", "error", err)
	}

	slog.Info("Server exiting")

}
