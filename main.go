package main

import (
	"takedi/xApi/configs"
	"takedi/xApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDb()

	routes.UserRoute(router)

	router.Run(":3000")
}
