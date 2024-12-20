package routes

import (
	"blog-plat/config"
	"blog-plat/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine, cfg config.Config) {
	services.InitDB(cfg.DatabaseURL)

	server.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"details":"OK!",
		})
	})
}