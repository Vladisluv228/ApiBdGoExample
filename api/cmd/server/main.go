package main

import (
	"log"

	"github.com/Vladisluv228/ApiBdGoExample/api/internal/db"
	"github.com/Vladisluv228/ApiBdGoExample/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Printf("Error initializing db: %v", err)
	}
	defer db.CloseDB()

	r := initServer()

	r.GET("/logs", handlers.GetUsersHandler)
	r.POST("/logs", handlers.CreateUserHandler)
	r.Run()
}

func initServer() *gin.Engine {
	r := gin.Default()
	return r
}
