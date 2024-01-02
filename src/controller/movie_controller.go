package movie_controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAllMovie(*gin.Context)
	GetMovieById(*gin.Context)
	CreateMovie(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
}

type Movie struct {
	ID          int64   `json:"id"`
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

type movieController struct {
	db *sql.DB
}

func NewMovieController(db *sql.DB) MovieController {
	return movieController{db: db}
}

func (m movieController) GetAllMovie(c *gin.Context) {

	getAll := "SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb"

	result, err := m.db.Query(getAll)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	defer result.Close()
	data := []Movie{}

	for result.Next() {
		var movie Movie
		if err := result.Scan(&movie.ID, &movie.ImdbID, &movie.Title, &movie.Year, &movie.Rating, &movie.IsSuperHero); err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, movie)
	}
	c.JSON(http.StatusOK, data)
}

func (m movieController) GetMovieById(c *gin.Context) {
	id := c.Param("id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	getById := "SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb WHERE id = ?"
	result := m.db.QueryRow(getById, movieID)

	var movie Movie
	err = result.Scan(&movie.ID, &movie.ImdbID, &movie.Title, &movie.Year, &movie.Rating, &movie.IsSuperHero)
	switch err {
	case nil:
		c.JSON(http.StatusOK, movie)
		return
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, map[string]string{"message!": "not found"})
		return
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (m movieController) CreateMovie(c *gin.Context) {
	movie := &Movie{}
	errorRequest := c.BindJSON(&movie)

	if errorRequest != nil {
		c.JSON(http.StatusBadRequest, errorRequest.Error())
		return
	}

	create, err := m.db.Prepare("INSERT INTO goimdb (imdbID,title,year,rating,isSuperHero) VALUES (?,?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, createError := create.Exec(movie.ImdbID, movie.Title, movie.Year, movie.Rating, movie.IsSuperHero)
	if createError != nil {
		c.JSON(500, createError)
		return
	}

	defer create.Close()
	id, _ := result.LastInsertId()
	movie.ID = id

	c.JSON(201, movie)
}

func (m movieController) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	var updates []string
	var values []interface{}
	for k, v := range data {
		updates = append(updates, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}
	values = append(values, movieID)
	query := fmt.Sprintf("UPDATE goimdb SET %s WHERE id = ?", strings.Join(updates, ", "))
	stmt, err := m.db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare update statement"})
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(values...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute update query"})
		return
	}

	numRows, _ := result.RowsAffected()
	if numRows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func (m movieController) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	delete, _ := m.db.Prepare("DELETE FROM goimdb where id = ? ")
	result, deleteError := delete.Exec(movieID)
	if deleteError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteError.Error()})
		return
	}

	numRows, _ := result.RowsAffected()
	if numRows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})

}
