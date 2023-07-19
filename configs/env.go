package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("MONGOURI"))
	return os.Getenv("MONGOURI")
}

func AdminPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("adminPassword"))
	return os.Getenv("adminPassword")
}

func AdminEmail() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("adminEmail"))
	return os.Getenv("adminEmail")
}
