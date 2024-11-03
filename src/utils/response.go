package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/razidev/movie-festival/src/models"
)

func MovieResponse(movie models.Movies) interface{} {
	return gin.H{
		"unique_id":   movie.UniqueID,
		"title":       movie.Title,
		"description": movie.Description,
		"duration":    movie.Duration,
		"artist":      movie.Artists,
		"genre_ids":   movie.GenreIds,
		"watch_url":   movie.WatchUrl,
		"voters":      movie.Voters,
		"viewers":     movie.Viewers,
	}
}
