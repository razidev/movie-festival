package routers

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	var validate = validator.New()

	movieGroup := r.Group("/movie")
	movieRoute(movieGroup, validate)

	return r
}
