package repository

import (
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type GenreRepository interface {
	ListGenres(search []uint) ([]models.Genres, error)
	UpdateViewers(ids []uint, genres []models.Genres) error
	HighestViewer() (models.Genres, error)
}

type genreRepository struct{}

func NewGenreRepository() GenreRepository {
	return &genreRepository{}
}

func (r *genreRepository) ListGenres(ids []uint) ([]models.Genres, error) {
	var genres []models.Genres

	if err := configs.DB.Where("id IN ?", ids).Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *genreRepository) UpdateViewers(ids []uint, genres []models.Genres) error {
	if err := configs.DB.Model(&models.Genres{}).Where("id IN ?", ids).
		Save(&genres).Error; err != nil {
		return err
	}

	return nil
}

func (r *genreRepository) HighestViewer() (models.Genres, error) {
	var genre models.Genres
	if err := configs.DB.Order("viewers DESC").First(&genre).Error; err != nil {
		return genre, err
	}

	return genre, nil
}
