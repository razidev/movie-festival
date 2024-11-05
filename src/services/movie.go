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
	FindHighestViewers() (map[string]interface{}, error)
	ListGenres() ([]models.Genres, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
	genreRepository repository.GenreRepository
}

func NewMovieService(movieRepo repository.MovieRepository, genreRepo repository.GenreRepository) MovieService {
	return &movieService{
		movieRepository: movieRepo,
		genreRepository: genreRepo,
	}
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

func (s *movieService) FindHighestViewers() (map[string]interface{}, error) {
	movie, err := s.movieRepository.HighestScore("viewers")
	if err != nil {
		return nil, err
	}

	genre, err := s.genreRepository.HighestViewer()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Movie": map[string]interface{}{
			"title":     movie.Title,
			"duration":  movie.Duration,
			"artists":   movie.Artists,
			"watch_url": movie.WatchUrl,
			"viewers":   movie.Viewers,
		},
		"Genre": map[string]interface{}{
			"name":    genre.Name,
			"viewers": genre.Viewers,
		},
	}, nil
}

func (s *movieService) ListGenres() ([]models.Genres, error) {
	genres, err := s.genreRepository.AllGenres()
	if err != nil {
		return genres, errors.New("Failed to list genres")
	}

	return genres, nil
}
