package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	exception "github.com/razidev/movie-festival/src/exceptions"
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
	search := ctx.Query("search")
	if errPage != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	if errLimit != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	movies, err := ctrl.Service.ListMovie(limit, page, search)
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

func (ctrl *UserController) PutWatchMovie(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")

	movie, err := ctrl.Service.UpdateViewers(uuid.MustParse(uniqueId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Watching movies...",
		"data": gin.H{
			"movie": utils.MovieResponse(movie),
		},
	})
}

func (ctrl *UserController) PostCreateUser(ctx *gin.Context) {
	var payload utils.PayloadUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	_, isExist := ctrl.Service.FindByEmail(payload.Email)
	if isExist {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	newUser, err := ctrl.Service.CreateUser(payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"user": utils.UserResponse(newUser),
		},
	})
}

func (ctrl *UserController) PostLoginUser(ctx *gin.Context) {
	var payload utils.PayloadUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := ctrl.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	_, isExist := ctrl.Service.FindByEmail(payload.Email)
	if !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := ctrl.Service.LoginUser(payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User login successfully",
		"data": gin.H{
			"token": token,
		},
	})
}

func (ctrl *UserController) PutVotesMovie(ctx *gin.Context) {
	movieUniqueId := ctx.Param("uniqueId")
	claims, exists := ctx.Get("unique_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User unauthorized"})
		return
	}
	claimsJSON, _ := json.Marshal(claims)
	userUniqueId := string(claimsJSON)

	err := ctrl.Service.VoteMovie(uuid.MustParse(movieUniqueId), uuid.MustParse(userUniqueId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Vote Movie successfully",
	})
}

func (ctrl *UserController) PutUnVotesMovie(ctx *gin.Context) {
	movieUniqueId := ctx.Param("uniqueId")
	claims, exists := ctx.Get("unique_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User unauthorized"})
		return
	}
	claimsJSON, _ := json.Marshal(claims)
	userUniqueId := string(claimsJSON)

	err := ctrl.Service.UnVoteMovie(uuid.MustParse(movieUniqueId), uuid.MustParse(userUniqueId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Unvote Movie successfully",
	})
}

func (ctrl *UserController) GetUserVotes(ctx *gin.Context) {
	users, err := ctrl.Service.ListUserVotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Users vote movies",
		"data":    utils.UserVotedResponse(users),
	})
}
