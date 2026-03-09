package main

import (
	"user-service/internal/app"
	"user-service/internal/bootstrap"
	"user-service/internal/config"
	internalHttp "user-service/internal/http"
	"user-service/internal/infrastructure/db"
	"user-service/internal/infrastructure/rabbitmq"
)

func main() {
	cfg := config.LoadConfig()

	database := db.GetPostgresDB()

	rmq := rabbitmq.NewRabbitMQ(cfg.RabbitMQURL)
	defer rmq.Close()

	container := app.NewContainer(database)

	bootstrap.StartConsumers(rmq, container)

	router := internalHttp.SetupRouter(cfg, container)

	srv := bootstrap.NewHTTPServer(router, cfg.Port, cfg.ReadTimeout, cfg.WriteTimeout, cfg.IdleTimeout)

	bootstrap.StartHTTPServer(srv, cfg)

	// --- Graceful shutdown ---
	bootstrap.WaitForShutdown(srv, cfg)
}
