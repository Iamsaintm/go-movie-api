package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func Database() *gorm.DB {
	databaseString := os.Getenv("DATABASE")

	db, err := gorm.Open(mysql.Open(databaseString), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to Database")
	return db
}
