package main

import (
	"log"

	"github.com/Vladisluv228/ApiBdGoExample/app/internal/db"
	"github.com/Vladisluv228/ApiBdGoExample/app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Printf("Error initializing db: %v", err)
	}
	defer db.CloseDB()

	r := initServer()

	r.GET("/logs", handlers.GetLogsHandler)
	r.POST("/logs", handlers.CreateLogHandler)
	r.Run()
}

func initServer() *gin.Engine {
	r := gin.Default()
	return r
}
