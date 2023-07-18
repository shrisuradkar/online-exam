package middleware

import (
	"log"
	"net/http"
	"onlineExam/helpers"
	"onlineExam/responses"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "No Authorization header provided"}})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err}})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		log.Println("In auth fucntion printing user_type")

		log.Println(c.Get("user_type"))

		c.Next()
	}

}
