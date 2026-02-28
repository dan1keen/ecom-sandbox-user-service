package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user-service/db"
	"user-service/internal/app"

	"user-service/config"
	internalHttp "user-service/internal/http"
	"user-service/server"
)

func main() {
	cfg := config.LoadConfig()

	database := db.GetPostgresDB()

	container := app.NewContainer(cfg, database)

	router := internalHttp.SetupRouter(cfg, container)

	srv := server.NewHTTPServer(router, cfg.Port, cfg.ReadTimeout, cfg.WriteTimeout, cfg.IdleTimeout)

	startServer(srv, cfg)

	// --- Graceful shutdown ---
	waitForShutdown(srv, cfg)
}

func startServer(srv *http.Server, cfg *config.Config) {
	go func() {
		log.Printf("Server running on port %s (%s mode)", cfg.Port, cfg.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
}

func waitForShutdown(srv *http.Server, cfg *config.Config) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
