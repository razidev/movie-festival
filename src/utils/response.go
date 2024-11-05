package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/models"
	"gorm.io/datatypes"
)

type Movie struct {
	UniqueId    uuid.UUID      `json:"unique_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int64          `json:"duration"`
	Artists     datatypes.JSON `json:"artists"`
	GenreIds    datatypes.JSON `json:"genre_ids"`
	WatchUrl    string         `json:"watch_url"`
	Voters      int            `json:"voters"`
	Viewers     int            `json:"viewers"`
}

func MovieResponse(movie models.Movies) Movie {
	return Movie{
		UniqueId:    movie.UniqueID,
		Title:       movie.Title,
		Description: movie.Description,
		Duration:    movie.Duration,
		Artists:     movie.Artists,
		GenreIds:    movie.GenreIds,
		WatchUrl:    movie.WatchUrl,
		Voters:      movie.Voters,
		Viewers:     movie.Viewers,
	}
}

func MoviesResponse(movies []models.Movies) []Movie {
	var responses []Movie
	for _, movie := range movies {
		responses = append(responses, MovieResponse(movie))
	}

	return responses
}

type User struct {
	UniqueId  uuid.UUID `json:"unique_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func UserResponse(user models.User) User {
	return User{
		UniqueId:  user.UniqueID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func UserVotedResponse(users []models.User) []string {
	var responses []string

	for _, user := range users {
		responses = append(responses, user.Email)
	}

	return responses
}

type Genres struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ListGenresResponse(genres []models.Genres) (restructure []Genres) {
	for _, genre := range genres {
		genre := Genres{
			ID:   genre.ID,
			Name: genre.Name,
		}

		restructure = append(restructure, genre)
	}

	return restructure
}
