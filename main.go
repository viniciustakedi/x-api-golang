package main

import (
	"net/http"
	"takedi/xApi/api/routes"
	"takedi/xApi/configs"
	"takedi/xApi/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	database.Connect()

	routes.UserRoute(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Working!",
		})
	})

	router.Run(":" + configs.EnvPort())
}
