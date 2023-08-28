package controller

import (
	"onlineExam/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ExamCollection *mongo.Collection = configs.GetCollection(configs.DB, "exam")

func GetAllExam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func DeleteExam() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func EditExam() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetExamById() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func CreateExam() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
