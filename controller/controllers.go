package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"onlineExam/configs"
	"onlineExam/helpers"
	"onlineExam/models"
	"onlineExam/responses"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userMgmtCollection *mongo.Collection = configs.GetCollection(configs.DB, "userMgmt")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			log.Fatal("err", err)
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			log.Fatal("Validation Error", validationErr)
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		count, err := userMgmtCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "error occured while checking for the email"}})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userMgmtCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "error occured while checking for the phone number"}})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "This email or phone number already exists"}})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.Course, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionCount, insertErr := userMgmtCollection.InsertOne(ctx, user)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "User item was not created"}})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, responses.ExamResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": resultInsertionCount}})
	}
}

func GetUserByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("---")
		log.Println(c.Get("user_type"))
		// if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userType := strings.ToUpper(c.Param("userType")) // c.Param("user_type")
		log.Println("-", userType)
		defer cancel()

		var user models.User
		if userType == "ADMIN" {
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": "Admin details are restricred"}})
			return
		}
		count, err := userMgmtCollection.EstimatedDocumentCount(ctx)
		log.Println(count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if count <= 0 {
			log.Println("Count is less than 0")

			c.JSON(http.StatusNotFound, responses.ExamResponse{Status: http.StatusNotFound, Message: "Error", Data: map[string]interface{}{"data": "Not a Singel User Present in Database"}})
			return
		}

		filter := bson.D{{"userType", userType}}
		result, err := userMgmtCollection.Find(ctx, filter)
		result.Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ExamResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": user}})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := strings.ToUpper(c.Param("userId")) // c.Param("user_type")
		log.Println("-", userId)
		defer cancel()

		var user models.User
		err := userMgmtCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ExamResponse{Status: http.StatusNotFound, Message: "Error", Data: map[string]interface{}{"data": "User not present in Database"}})
			return
		}

		if *user.User_type == "ADMIN" {
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": "Admin can't be Delete"}})
			return
		}

		result, err := userMgmtCollection.DeleteOne(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ExamResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": "User Deleted Successfully"}})
	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var founduser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		err := userMgmtCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode((&founduser))
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "Email ID is incorrect"}})
			return
		}

		passwordIsValid, msg := VerifyPasword(*user.Password, *founduser.Password)
		defer cancel()

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": msg}})
			return
		}

		if founduser.Email == nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": "User Not Found"}})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(*founduser.Email, *founduser.First_name, *founduser.Last_name, *founduser.Course, *founduser.User_type, *&founduser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, founduser.User_id)
		err = userMgmtCollection.FindOne(ctx, bson.M{"user_id": founduser.User_id}).Decode(&founduser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, founduser)

	}
}
func VerifyPasword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Password is Incorrect")
		check = false
	}
	return check, msg
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := strings.ToUpper(c.Param("userId")) // c.Param("user_type")
		log.Println("-", userId)
		defer cancel()

		var user models.User
		err := userMgmtCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ExamResponse{Status: http.StatusNotFound, Message: "Error", Data: map[string]interface{}{"data": "User not present in Database"}})
			return
		}

		if user.User_id == "ADMIN" {
			c.JSON(http.StatusBadRequest, responses.ExamResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": "Admin can't be Update"}})
			return
		}

		filter := bson.M{"user_id": userId}
		update := bson.M{
			"first_name": user.First_name,
			"last_name":  user.Last_name,
			"Password":   user.Password,
			"email":      user.Email,
			"phone":      user.Phone,
			"course":     user.Course,
		}
		result, err := userMgmtCollection.UpdateOne(ctx, filter, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedUser models.User
		if result.MatchedCount == 1 {
			if err := userMgmtCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode((&updatedUser)); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, responses.ExamResponse{Status: http.StatusInternalServerError, Message: "NOT UPDATED", Data: map[string]interface{}{"data": "User is not updated"}})
			return
		}

		c.JSON(http.StatusOK, responses.ExamResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": updatedUser}})
	}
}
