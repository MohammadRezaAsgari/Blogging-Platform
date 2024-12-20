package main

import (
	"blog-plat/config"
	"blog-plat/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	server := gin.Default()

	routes.SetupRoutes(server, cfg)
	server.Run(":8080")
}