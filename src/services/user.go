package services

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/repository"
)

type UserService interface {
	ListMovie(offest int, page int, search string) ([]models.Movies, error)
	UpdateViewers(unique uuid.UUID) (models.Movies, error)
}

type userService struct {
	movieRepository repository.MovieRepository
	genreRepository repository.GenreRepository
}

func NewUserService(movieRepo repository.MovieRepository, genreRepo repository.GenreRepository) UserService {
	return &userService{
		movieRepository: movieRepo,
		genreRepository: genreRepo,
	}
}

func (s *userService) ListMovie(limit int, page int, search string) ([]models.Movies, error) {
	offset := (page - 1) * limit

	movies, err := s.movieRepository.ListMovies(offset, limit, search)
	if err != nil {
		return movies, errors.New("Movies not found")
	}

	return movies, nil
}

func (s *userService) UpdateViewers(uniqueID uuid.UUID) (models.Movies, error) {
	foundMovie, err := s.movieRepository.GetMovieByUniqueId(uniqueID)
	if err != nil {
		return foundMovie, errors.New("Movie not found")
	}

	foundMovie.Viewers++
	_, errUpdate := s.movieRepository.UpdateMovie(foundMovie)
	if errUpdate != nil {
		return foundMovie, errors.New("failed to update movie")
	}

	var ids []uint
	if err := json.Unmarshal(foundMovie.GenreIds, &ids); err != nil {
		return foundMovie, errors.New("Error unmarshalling JSON")
	}

	genres, err := s.genreRepository.ListGenres(ids)
	for _, genre := range genres {
		genre.Viewers++
		genres = append(genres, genre)
	}

	err = s.genreRepository.UpdateViewers(ids, genres)
	if err != nil {
		return foundMovie, errors.New("Failed to update Genres")
	}

	return foundMovie, nil
}
