package routers

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/razidev/movie-festival/src/controllers"
	middleware "github.com/razidev/movie-festival/src/middlewares"
	"github.com/razidev/movie-festival/src/repository"
	"github.com/razidev/movie-festival/src/services"
)

func UserRoute(group *gin.RouterGroup, validator *validator.Validate) {
	movieRepo := repository.NewMovieRepository()
	genreRepo := repository.NewGenreRepository()
	userRepo := repository.NewUserRepository()
	userVoteRepo := repository.NewUserVoteRepository()

	userService := services.NewUserService(userRepo, movieRepo, genreRepo, userVoteRepo)
	userController := controllers.NewUserController(userService, validator)

	group.GET("/movies", userController.GetMovies)
	group.PUT("/movies/:uniqueId", userController.PutWatchMovie)
	group.POST("/register", userController.PostCreateUser)
	group.POST("/login", userController.PostLoginUser)

	group.Use(middleware.JWTAuthMiddleware())
	group.PUT("/movies/votes/:uniqueId", userController.PutVotesMovie)
	group.PUT("/movies/unvotes/:uniqueId", userController.PutUnVotesMovie)
}
