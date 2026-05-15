package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	assets "visualizationBdDebet"
	"visualizationBdDebet/internal/contractor"
	"visualizationBdDebet/internal/customer"
	"visualizationBdDebet/internal/object"
	"visualizationBdDebet/internal/response"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"visualizationBdDebet/internal/blockfactor"
	"visualizationBdDebet/internal/config"
	"visualizationBdDebet/internal/contract"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/delivery/router"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Application stopped with error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	db, err := sqlx.Connect("postgres", cfg.DBConnString())
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("Failed to close db", "error", err)
		}
	}()

	debetRepo := debet.NewRepository(db)
	debetService := debet.NewService(debetRepo)
	debetHandler := debet.NewHandler(debetService)

	contractRepo := contract.NewRepository(db)
	contractService := contract.NewService(contractRepo)
	contractHandler := contract.NewHandler(contractService)

	blockfactorRepo := blockfactor.NewRepository(db)
	blockfactorService := blockfactor.NewService(blockfactorRepo)
	blockfactorHandler := blockfactor.NewHandler(blockfactorService)

	responseRepo := response.NewRepository(db)
	responseService := response.NewService(responseRepo, debetService)
	responseHandler := response.NewHandler(responseService)

	customerRepo := customer.NewRepository(db)
	customerService := customer.NewService(customerRepo)
	customerHandler := customer.NewHandler(customerService)

	contractorRepo := contractor.NewRepository(db)
	contractorService := contractor.NewService(contractorRepo)
	contractorHandler := contractor.NewHandler(contractorService)

	objectsRepo := object.NewRepository(db)
	objectsService := object.NewService(objectsRepo)
	objectHandler := object.NewHandler(objectsService)

	apiRouter := router.NewRouter(
		debetHandler,
		contractHandler,
		blockfactorHandler,
		responseHandler,
		customerHandler,
		contractorHandler,
		objectHandler,
	)

	staticFS, err := fs.Sub(assets.FS, "web")
	if err != nil {
		return fmt.Errorf("get static files subdir: %w", err)
	}

	apiRouter.PathPrefix("/").Handler(http.FileServer(http.FS(staticFS)))

	srv := &http.Server{
		Handler:      apiRouter,
		Addr:         ":" + cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	serveErr := make(chan error, 1)

	go func() {
		slog.Info("starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case sig := <-quit:
		slog.Info("Shutdown signal received", "signal", sig.String())
	case err := <-serveErr:
		return fmt.Errorf("start server: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	slog.Info("Server exiting")
	return nil
}
