package main

import (
	"fmt"
	"os"

	router "github.com/Iamsaintm/go-movie-api/src/Router"
	"github.com/Iamsaintm/go-movie-api/src/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	engine := gin.Default()

	port := os.Getenv("PORT")

	db := initDatabase()

	server := config.Server{Engine: engine, DB: db}
	router.MovieRouter(server, "/movie")

	fmt.Println("Server running on port", port)
	server.Engine.Run(":" + port)
}

func initDatabase() *gorm.DB {
	databaseString := os.Getenv("DATABASE")

	db, err := gorm.Open(mysql.Open(databaseString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	fmt.Println("Connected to Database")
	return db
}
