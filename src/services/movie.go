package services

import (
	"errors"
	"net/url"

	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/repository"
)

type MovieService interface {
	CreateMovie(movie models.Movies) (models.Movies, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{movieRepository: repo}
}

func (s *movieService) CreateMovie(movie models.Movies) (models.Movies, error) {
	url, errUrl := url.Parse(movie.WatchUrl)
	if errUrl != nil || url.Host == "" || url.Scheme == "" {
		return movie, errors.New("Movie Url is not valid")
	}

	newMovie, err := s.movieRepository.CreateMovie(movie)
	if err != nil {
		return newMovie, err
	}

	return newMovie, nil
}
