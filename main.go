package main

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"visualizationBdDebet/internal/delivery/response"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"visualizationBdDebet/internal/blockfactor"
	"visualizationBdDebet/internal/config"
	"visualizationBdDebet/internal/contract"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/delivery/handler"
	"visualizationBdDebet/internal/delivery/router"
)

//go:embed web
var staticFiles embed.FS

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

	contractRepo := contract.NewRepository(db)
	contractService := contract.NewService(contractRepo)
	contractHandler := handler.NewContractHandler(contractService)

	blockfactorRepo := blockfactor.NewRepository(db)
	blockfactorService := blockfactor.NewService(blockfactorRepo)
	blockfactorHandler := handler.NewBlockFactorHandler(blockfactorService)

	responseRepo := response.NewRepository(db)
	responseService := response.NewService(responseRepo, debetService)
	responseHandler := response.NewResponseHandler(responseService)

	apiRouter := router.NewRouter(
		debetHandler,
		contractHandler,
		blockfactorHandler,
		responseHandler,
	)

	staticFS, err := fs.Sub(staticFiles, "web")
	if err != nil {
		slog.Error("Failed to get static files subdir", "error", err)
		panic(err)
	}

	apiRouter.PathPrefix("/").Handler(http.FileServer(http.FS(staticFS)))

	/*gateway := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Проверяем, относится ли путь к API
		if strings.HasPrefix(path, "/debet") ||
			strings.HasPrefix(path, "/contract") ||
			strings.HasPrefix(path, "/blockfactor") ||
			strings.HasPrefix(path, "/response") {
			apiRouter.ServeHTTP(w, r)
			return
		}
		// Иначе отдаём статику
		http.FileServer(http.FS(staticFS)).ServeHTTP(w, r)
	})*/

	srv := http.Server{
		Handler:      apiRouter,
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
