package main

import (
	"log"
	"net/http"
	"takedi/xApi/configs"
	"takedi/xApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	configs.ConnectDb()

	routes.UserRoute(router)

	router.GET("/ping", func(c *gin.Context) {
		log.Println("Ping route accessed")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := configs.EnvPort()

	if port == "" {
		port = "80"
	}

	log.Printf("Starting server on port %s\n", port)

	router.Run(":" + port)
}
