package routes

import (
	"onlineExam/controller"
	"onlineExam/middleware"

	"github.com/gin-gonic/gin"
)

func ExamRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.GET("/getAllExam/", controller.GetAllExam())
	router.DELETE("/deleteExam/:examId", controller.DeleteExam())
	router.PUT("/editExam/:examId", controller.EditExam())
	router.PUT("/getExam/:examId", controller.GetExamById())
	router.POST("/createExam/:examId", controller.CreateExam())
}
