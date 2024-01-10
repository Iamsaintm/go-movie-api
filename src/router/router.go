package router

import (
	"github.com/Iamsaintm/go-movie-api/src/config"
	movie_controller "github.com/Iamsaintm/go-movie-api/src/controller"
)

func MovieRouter(server config.Server, PATH string) {

	movieRouter := server.Engine.Group(PATH)

	db := config.Database()

	movieController := movie_controller.NewMovieController(db)

	movieRouter.GET("/getAll", movieController.GetAllMovie)
	movieRouter.GET("/:id", movieController.GetMovieById)
	movieRouter.POST("/create", movieController.CreateMovie)
	movieRouter.PATCH("/update/:id", movieController.UpdateMovie)
	movieRouter.DELETE("/delete/:id", movieController.DeleteMovie)
}
