package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/razidev/movie-festival/src/services"
	"github.com/razidev/movie-festival/src/utils"
)

type UserController struct {
	Service  services.UserService
	Validate *validator.Validate
}

func NewUserController(service services.UserService, validate *validator.Validate) *UserController {
	return &UserController{
		Service:  service,
		Validate: validate,
	}
}

func (ctrl *UserController) GetMovies(ctx *gin.Context) {
	page, errPage := strconv.Atoi(ctx.Query("page"))
	limit, errLimit := strconv.Atoi(ctx.Query("limit"))
	if errPage != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	if errLimit != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	movies, err := ctrl.Service.ListMovie(limit, page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"movies": utils.MoviesResponse(movies),
		},
	})
}
