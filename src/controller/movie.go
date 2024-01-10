package movie_controller

import (
	"net/http"

	"github.com/Iamsaintm/go-movie-api/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController interface {
	GetAllMovie(*gin.Context)
	GetMovieById(*gin.Context)
	CreateMovie(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
}

type movieController struct {
	db *gorm.DB
}

func NewMovieController(db *gorm.DB) MovieController {
	return &movieController{db: db}
}

func (m *movieController) GetAllMovie(c *gin.Context) {
	var movies []model.Movie
	m.db.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func (m *movieController) GetMovieById(c *gin.Context) {
	id := c.Param("id")
	var movie model.Movie
	if err := m.db.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *movieController) CreateMovie(c *gin.Context) {
	movie := &model.Movie{}
	if err := c.BindJSON(movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.db.Create(movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (m *movieController) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie model.Movie
	if err := m.db.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.db.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func (m *movieController) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	var movie model.Movie
	if err := m.db.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := m.db.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
