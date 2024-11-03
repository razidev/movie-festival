package repository

import (
	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type MovieRepository interface {
	CreateMovie(movie models.Movies) (models.Movies, error)
}

type movieRepository struct {
}

func NewMovieRepository() MovieRepository {
	return &movieRepository{}
}

func (r *movieRepository) CreateMovie(movie models.Movies) (models.Movies, error) {
	movie.UniqueID = uuid.New()
	if err := configs.DB.Create(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}
