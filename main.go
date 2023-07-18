package main

import (
	"onlineExam/configs"
	"onlineExam/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	configs.ConnectDB()

	routes.AuthRoutes(server)
	routes.UserMgmtRoutes(server)

	server.Run(":9000")
}
