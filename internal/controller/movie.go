package controller

import (
	"net/http"

	"github.com/GareArc/MovieMate/internal/service"
	"github.com/gin-gonic/gin"
)

type MovieController struct{}

func (mc *MovieController) GetMovieInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "field id is reqired",
		})
		return
	}

	movie_service := service.MovieService{}
	movie, err := movie_service.GetMovieInfoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error during processing the movie",
		})
		return
}

	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) GetShowTimes(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "field id is required",
		})
		return
	}

	movie_service := service.MovieService{}
	showtime := movie_service.GetShowTimeListById(id)

	c.JSON(http.StatusOK, showtime)
}
