package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	middleware "github.com/razidev/movie-festival/src/middlewares"
	"github.com/razidev/movie-festival/src/models"
	"github.com/razidev/movie-festival/src/repository"
)

type UserService interface {
	ListMovie(offest int, page int, search string) ([]models.Movies, error)
	UpdateViewers(unique uuid.UUID) (models.Movies, error)
	FindByEmail(email string) (models.User, bool)
	CreateUser(email string, password string) (models.User, error)
	LoginUser(email string, password string) (string, error)
	VoteMovie(movieUniueId uuid.UUID, userUniqueId uuid.UUID) error
	UnVoteMovie(movieUniueId uuid.UUID, userUniqueId uuid.UUID) error
}

type userService struct {
	userRepository     repository.UserRepository
	movieRepository    repository.MovieRepository
	genreRepository    repository.GenreRepository
	userVoteRepository repository.UserVoteRepository
}

func NewUserService(userRepo repository.UserRepository, movieRepo repository.MovieRepository, genreRepo repository.GenreRepository, userVoteRepo repository.UserVoteRepository) UserService {
	return &userService{
		userRepository:     userRepo,
		movieRepository:    movieRepo,
		genreRepository:    genreRepo,
		userVoteRepository: userVoteRepo,
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

func (s *userService) FindByEmail(email string) (models.User, bool) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, false
	}

	return user, true
}

func (s *userService) CreateUser(email string, password string) (models.User, error) {
	var user models.User

	hashedPass, err := middleware.HashPassword(password)
	if err != nil {
		return user, errors.New("Failed to hash password")
	}
	user.Email = email
	user.Password = hashedPass

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return user, errors.New("Failed to create user")
	}
	return newUser, nil
}

func (s *userService) LoginUser(email string, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("User not found")
	}

	if !middleware.CheckPasswordHash(password, user.Password) {
		return "", errors.New("Invalid username or password")
	}

	token, err := middleware.GenerateJWT(user.Email, user.UniqueID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *userService) VoteMovie(movieUniqueId uuid.UUID, userUniqueId uuid.UUID) error {
	foundMovie, err := s.movieRepository.GetMovieByUniqueId(movieUniqueId)
	if err != nil {
		return errors.New("Movie not found")
	}

	foundMovie.Voters++
	_, err = s.movieRepository.UpdateMovie(foundMovie)
	if err != nil {
		return errors.New("failed to update movie")
	}

	_, err = s.userVoteRepository.FindCurrentVote(movieUniqueId, userUniqueId)
	if err == nil {
		return errors.New("User has already voted for this movie")
	}

	newVote := models.UserVotes{
		MovieUniqueID: movieUniqueId,
		UserUniqueID:  userUniqueId,
		Status:        "voted",
	}

	err = s.userVoteRepository.CreateVote(newVote)
	if err != nil {
		return errors.New("Failed to create user vote")
	}

	return nil
}

func (s *userService) UnVoteMovie(movieUniqueId uuid.UUID, userUniqueId uuid.UUID) error {
	vote, err := s.userVoteRepository.FindCurrentVote(movieUniqueId, userUniqueId)
	fmt.Println("vote", vote)
	if err != nil {
		return errors.New("User Vote not found")
	}

	if vote.Status == "unvoted" {
		return errors.New("User has not voted for this movie")
	}

	foundMovie, err := s.movieRepository.GetMovieByUniqueId(movieUniqueId)
	if err != nil {
		return errors.New("Movie not found")
	}

	foundMovie.Voters--
	_, err = s.movieRepository.UpdateMovie(foundMovie)
	if err != nil {
		return errors.New("failed to update movie")
	}

	err = s.userVoteRepository.UpdateVote(vote.ID, "unvoted")
	if err != nil {
		return errors.New("Failed to unvote movie")
	}

	return nil
}
