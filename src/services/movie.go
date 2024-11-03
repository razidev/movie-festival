package services

import (
	"errors"

	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/repository"
)

type MovieService interface {
	CreateMovie(movie models.Movies) (models.Movies, error)
	UpdateMovie(movie models.Movies) (models.Movies, error)
	FindHighestVotes() (models.Movies, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{movieRepository: repo}
}

func (s *movieService) CreateMovie(movie models.Movies) (models.Movies, error) {
	newMovie, err := s.movieRepository.CreateMovie(movie)
	if err != nil {
		return newMovie, err
	}

	return newMovie, nil
}

func (s *movieService) UpdateMovie(movie models.Movies) (models.Movies, error) {
	foundMovie, err := s.movieRepository.GetMovieByUniqueId(movie.UniqueID)
	if err != nil {
		return foundMovie, errors.New("Movie not found")
	}

	foundMovie.Title = movie.Title
	foundMovie.Description = movie.Description
	foundMovie.Duration = movie.Duration
	foundMovie.Artists = movie.Artists
	foundMovie.GenreIds = movie.GenreIds
	foundMovie.WatchUrl = movie.WatchUrl

	updateMovie, err := s.movieRepository.UpdateMovie(foundMovie)
	if err != nil {
		return updateMovie, errors.New("failed to update movie")
	}

	return updateMovie, nil
}

func (s *movieService) FindHighestVotes() (models.Movies, error) {
	movie, err := s.movieRepository.HighestScore("voters")
	if err != nil {
		return movie, errors.New("Highest voter not found")
	}

	return movie, nil
}
