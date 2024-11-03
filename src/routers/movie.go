package routers

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/razidev/movie-festival/src/controllers"
	"github.com/razidev/movie-festival/src/repository"
	"github.com/razidev/movie-festival/src/services"
)

func movieRoute(group *gin.RouterGroup, validator *validator.Validate) {
	movieRepo := repository.NewMovieRepository()
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMoviesController(movieService, validator)

	group.POST("/", movieController.PostMovie)
	group.PUT("/:uniqueId", movieController.PutMovie)
	group.GET("/highest-vote", movieController.GetHighestVotes)
}
