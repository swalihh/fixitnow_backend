package main

import (
	"log"
	"service-api/database"
	"service-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	database.ConnectToDatabase()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	routes.ServiceRouter(server)
	routes.AdminRoutes(server)
	routes.UserRoutes(server)
	log.Println("Server started at 5000 port")
	server.Run(":5000")
}
