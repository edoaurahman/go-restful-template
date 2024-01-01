package main

import (
	"log"

	"github.com/edoaurahman/go-restful-template/src/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}
	server.Init()
}
