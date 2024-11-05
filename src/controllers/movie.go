package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	exception "github.com/razidev/movie-festival/src/exceptions"
	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/services"
	"github.com/razidev/movie-festival/src/utils"
	"gorm.io/datatypes"
)

type MoviesController struct {
	Service  services.MovieService
	Validate *validator.Validate
}

func NewMoviesController(service services.MovieService, validate *validator.Validate) *MoviesController {
	return &MoviesController{
		Service:  service,
		Validate: validate,
	}
}

func (ctrl *MoviesController) PostMovie(ctx *gin.Context) {
	var payload utils.PayloadMovie
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	ArtistName, _ := json.Marshal(payload.ArtistName)
	GenreIds, _ := json.Marshal(payload.GenreIds)

	movie := models.Movies{
		Title:       payload.Title,
		Description: payload.Description,
		Duration:    payload.Duration,
		Artists:     datatypes.JSON(ArtistName),
		GenreIds:    datatypes.JSON(GenreIds),
		WatchUrl:    payload.WatchUrl,
	}

	newMovie, err := ctrl.Service.CreateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Movie created successfully",
		"data": gin.H{
			"movie": utils.MovieResponse(newMovie),
		},
	})
}

func (ctrl *MoviesController) PutMovie(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	var payload utils.PayloadMovie
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	ArtistName, _ := json.Marshal(payload.ArtistName)
	GenreIds, _ := json.Marshal(payload.GenreIds)

	movie := models.Movies{
		UniqueID:    uuid.MustParse(uniqueId),
		Title:       payload.Title,
		Description: payload.Description,
		Duration:    payload.Duration,
		Artists:     datatypes.JSON(ArtistName),
		GenreIds:    datatypes.JSON(GenreIds),
		WatchUrl:    payload.WatchUrl,
	}

	updatedMovie, err := ctrl.Service.UpdateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Movie updated successfully",
		"data": gin.H{
			"movie": utils.MovieResponse(updatedMovie),
		},
	})
}

func (ctrl *MoviesController) GetHighestVotes(ctx *gin.Context) {
	highestVotesMovie, err := ctrl.Service.FindHighestVotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Top 1 highest voted movies",
		"data": gin.H{
			"movie": utils.MovieResponse(highestVotesMovie),
		},
	})
}

func (ctrl *MoviesController) GetHighestViewers(ctx *gin.Context) {
	highestViewers, _ := ctrl.Service.FindHighestViewers()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Top 1 highest viewers movies",
		"data": gin.H{
			"movie": highestViewers["Movie"],
			"genre": highestViewers["Genre"],
		},
	})
}

func (ctrl *MoviesController) GetMovieGenres(ctx *gin.Context) {
	genres, err := ctrl.Service.ListGenres()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of genres",
		"data":    utils.ListGenresResponse(genres),
	})
}
