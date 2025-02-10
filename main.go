package main

import (
	"log"
	"petplate-auth/config"
	"petplate-auth/models"
	"petplate-auth/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})


	r := routes.SetupRouter()
	log.Println("Auth Service running on port 8081...")
	r.Run(":8081")
}
