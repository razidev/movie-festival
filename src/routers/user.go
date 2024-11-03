package routers

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/razidev/movie-festival/src/controllers"
	"github.com/razidev/movie-festival/src/repository"
	"github.com/razidev/movie-festival/src/services"
)

func UserRoute(group *gin.RouterGroup, validator *validator.Validate) {
	movieRepo := repository.NewMovieRepository()
	userService := services.NewUserService(movieRepo)
	userController := controllers.NewUserController(userService, validator)

	group.GET("/movies", userController.GetMovies)
}
