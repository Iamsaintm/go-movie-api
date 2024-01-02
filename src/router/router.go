package router

import (
	movie_controller "github.com/Iamsaintm/go-movie-api/src/controller"
	"github.com/Iamsaintm/go-movie-api/src/model"
	"github.com/Iamsaintm/go-movie-api/src/repo"
)

func MovieRouter(server model.Server, PATH string) {

	movieRouter := server.Engine.Group(PATH)

	db := repo.Database()

	movieController := movie_controller.NewMovieController(db)

	movieRouter.GET("/getAll", movieController.GetAllMovie)
	movieRouter.GET("/:id", movieController.GetMovieById)
	movieRouter.POST("/create", movieController.CreateMovie)
	movieRouter.PATCH("/update/:id", movieController.UpdateMovie)
	movieRouter.DELETE("/delete/:id", movieController.DeleteMovie)
}
