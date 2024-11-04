package services

import (
	"errors"

	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/repository"
)

type UserService interface {
	ListMovie(offest int, page int, search string) ([]models.Movies, error)
}

type userService struct {
	movieRepository repository.MovieRepository
}

func NewUserService(repo repository.MovieRepository) UserService {
	return &userService{movieRepository: repo}
}

func (s *userService) ListMovie(limit int, page int, search string) ([]models.Movies, error) {
	offset := (page - 1) * limit

	movies, err := s.movieRepository.ListMovies(offset, limit, search)
	if err != nil {
		return movies, errors.New("Movies not found")
	}

	return movies, nil
}
