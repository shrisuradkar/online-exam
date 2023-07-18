package routes

import (
	"onlineExam/controller"

	"github.com/gin-gonic/gin"
)

func UserMgmtRoutes(router *gin.Engine) {
	router.POST("/registerUser", controller.RegisterUser())
	router.POST("/login", controller.Login())

	// Admin can view both instructors and candidates
	// Instructor can only able to view candidates
}
