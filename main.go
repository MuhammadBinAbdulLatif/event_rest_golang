package main

import (
	"rest-api/db"
	"rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  db.ConnectDB()
  db.CreateTables()
  r := gin.Default()

  handlers.RegisterRoutes(r)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run()
}