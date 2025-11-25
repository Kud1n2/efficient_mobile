// @title Effective Mobile Subscriptions API
// @version 1.0
// @description Comprehensive REST API for managing subscriptions with CRUD operations and cost calculations

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@effectivemobile.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverPort := os.Getenv("SERVER_PORT")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}
	log.SetOutput(file)

	db, err = initDB()
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	router := gin.Default()
	router.GET("/effective_mobile", getAll)
	router.POST("/effective_mobile", post)
	router.GET("effective_mobile/:service_name", getByName)
	router.GET("getSum", getSum)
	router.DELETE("effective_mobile/id", deleteByServiceName)
	router.PUT("effective_mobile/id", updateByServiceName)

	router.Run(":" + serverPort)
	defer db.Close()
}
