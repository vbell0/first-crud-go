package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vbell0/first-crud-go/src/controller/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	routes.InitRoutes(r.Group("/api"))

	err = r.Run()
	if err != nil {
		return
	}
}
