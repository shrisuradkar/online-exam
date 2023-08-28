package services

import (
	"context"
	"log"
	"onlineExam/configs"
	"onlineExam/controller"
	"onlineExam/helpers"
	"onlineExam/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userMgmtCollection *mongo.Collection = configs.GetCollection(configs.DB, "userMgmt")

func Admin() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User

	err := userMgmtCollection.FindOne(ctx, bson.M{"user_type": "ADMIN"}).Decode(&user)
	if err == nil {
		log.Panic("Admin user is already created")
		return
	}
	email := configs.AdminEmail()
	userType := "ADMIN"
	name := "admin"
	phone := "1234567890"
	adminPassword := configs.AdminPassword()
	password := controller.HashPassword(adminPassword)
	user.Password = &password
	user.Email = &email
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.User_type = &userType
	user.First_name = &name
	user.Last_name = &name
	user.Course = &name
	user.Phone = &phone
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.Course, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	_, insertErr := userMgmtCollection.InsertOne(ctx, user)

	if insertErr != nil {
		log.Panic("Admin user is not created")
		return
	}
	// log.Println(resultInsertionCount)
	defer cancel()
}
