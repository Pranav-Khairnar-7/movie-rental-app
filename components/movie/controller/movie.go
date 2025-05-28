package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/movie-rental-app/components/movie/service"
	custError "github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
	uuid "github.com/satori/go.uuid"
)

type MovieController struct {
	MovieService *service.MovieService
}

func NewMovieController(svc *service.MovieService) *MovieController {
	return &MovieController{MovieService: svc}
}

func (c *MovieController) GetMovies(ctx *gin.Context) {
	var (
		limit  *int
		offset *int
	)

	if l := ctx.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = &val
		}
	}

	if o := ctx.Query("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil {
			offset = &val
		}
	}

	var movies []model.Movie

	err := c.MovieService.GetMovies(ctx, &movies, limit, offset)
	if err != nil {
		var httpErr *custError.HTTPError
		if errors.As(err, &httpErr) {
			ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	var (
		movieID uuid.UUID
	)

	if l := ctx.Param("movieID"); l != "" {
		movieID = uuid.FromStringOrNil(l)
	}

	var movie model.Movie

	err := c.MovieService.GetMovieByID(ctx, &movie, movieID)
	if err != nil {
		var httpErr *custError.HTTPError
		if errors.As(err, &httpErr) {
			ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, movie)
}
