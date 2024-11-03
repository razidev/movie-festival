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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorsMap})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Movie created successfully", "movie": newMovie})
}

func (ctrl *MoviesController) PutMovie(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	var payload utils.PayloadMovie
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorsMap})
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully", "movie": updatedMovie})
}
