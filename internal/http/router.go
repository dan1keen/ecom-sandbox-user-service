package http

import (
	"net/http"
	"user-service/internal/app"
	"user-service/internal/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, container *app.Container) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	registerHealthRoutes(router)
	registerApplicationRoutes(router, container)

	return router
}

func registerHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func registerApplicationRoutes(r *gin.Engine, container *app.Container) {
	r.GET("/profile", container.UserHandler().GetProfile())
}
