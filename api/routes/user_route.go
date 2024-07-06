package routes

import (
	userController "takedi/xApi/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user/register", userController.CreateUser())
	router.PATCH("/user/update/:userId", userController.UpdateUser())
	router.GET("/user/:userId", userController.GetUserById())
	router.DELETE("/user/:userId", userController.DeleteUser())
}
