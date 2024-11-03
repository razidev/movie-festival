package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type MovieRepository interface {
	CreateMovie(movie models.Movies) (models.Movies, error)
	GetMovieByUniqueId(unique uuid.UUID) (models.Movies, error)
	UpdateMovie(movie models.Movies) (models.Movies, error)
	HighestScore(column string) (models.Movies, error)
	ListMovies(offset int, limit int) ([]models.Movies, error)
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

func (r *movieRepository) GetMovieByUniqueId(uniqueId uuid.UUID) (models.Movies, error) {
	var movie models.Movies
	if err := configs.DB.Where("unique_id = ?", uniqueId).First(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}

func (r *movieRepository) UpdateMovie(movie models.Movies) (models.Movies, error) {
	if err := configs.DB.Save(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}

func (r *movieRepository) HighestScore(column string) (models.Movies, error) {
	query := fmt.Sprintf("%s DESC", column)
	var movie models.Movies
	if err := configs.DB.Order(query).First(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}

func (r *movieRepository) ListMovies(offset int, limit int) ([]models.Movies, error) {
	var movies []models.Movies
	if err := configs.DB.Offset(offset).Limit(limit).Find(&movies).Error; err != nil {
		return movies, err
	}

	return movies, nil
}
