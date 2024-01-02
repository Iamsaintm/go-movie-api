package main

import (
	"fmt"
	"os"

	"github.com/Iamsaintm/go-movie-api/src/model"
	"github.com/Iamsaintm/go-movie-api/src/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	engine := gin.Default()

	port := os.Getenv("PORT")

	r := model.Server{Engine: engine}
	router.MovieRouter(r, "/movie")
	fmt.Println("server run on port ", port)
	r.Engine.Run(":" + port)
}
