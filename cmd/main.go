package main

import (
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

	server.StartServer(srv, cfg)

	// --- Graceful shutdown ---
	server.WaitForShutdown(srv, cfg)
}
