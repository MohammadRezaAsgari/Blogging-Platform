package routes

import (
	"blog-plat/config"
	v1 "blog-plat/internal/api/v1"
	"blog-plat/internal/middlewares"
	"blog-plat/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine, cfg config.Config) {
	services.InitDB(cfg.DatabaseURL)

	server.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"details": "OK!",
		})
	})

	// Auth app
	server.POST("/api/v1/auth/login", v1.Login)
	server.POST("/api/v1/auth/register", v1.Register)
	server.GET("/api/v1/auth/me", middlewares.AuthRequired(), v1.GetUserProfile)

	// Blog app
	server.GET("/api/v1/blog/articles", v1.ArticleList)
	server.POST("/api/v1/blog/articles", middlewares.AuthRequired(), v1.CreateArticle)
	server.GET("/api/v1/blog/articles/:id", v1.ArticleByID)
	server.PATCH("/api/v1/blog/articles/:id", middlewares.AuthRequired(), v1.UpdateArticleByID)
	server.DELETE("/api/v1/blog/articles/:id", middlewares.AuthRequired(), v1.DeleteArticleByID)
}
