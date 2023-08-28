package routes

import (
	"onlineExam/controller"
	"onlineExam/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.GET("/getUserByType/:userType", controller.GetUserByType())
	router.DELETE("/deleteUser/:userId", controller.DeleteUser())
	router.PUT("/editUser/:userId", controller.EditUser())

}
