package main

import (
	"log"

	"github.com/WanKapef/go-api/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("API rodando na porta", cfg.Port)
	r.Run(":" + cfg.Port)
}
