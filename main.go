package main

import (
	"onlineExam/configs"
	"onlineExam/routes"
	"onlineExam/services"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	configs.ConnectDB()

	services.Admin()
	routes.UserMgmtRoutes(server)

	routes.AuthRoutes(server)
	routes.ExamRoutes(server)

	server.Run(":9000")
}
