package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/envde/effmobi/interanl/config"
	"github.com/envde/effmobi/interanl/pkg/logger"
	db "github.com/envde/effmobi/interanl/pkg/postgres"
	"github.com/envde/effmobi/interanl/transport"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config.Load", "err", err)
		os.Exit(1)
	}

	log := logger.New(cfg.App.Env)

	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg.DB.DSN())
	if err != nil {
		log.Error("db.NewPool", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	log.Info("connected to database")

	queries := db.New(pool)

	router := transport.NewRouter(queries, log)

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: router,
	}

	go func() {
		log.Info("server started", "port", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("ListenAndServe", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("srv.Shutdown", "err", err)
	}
}
