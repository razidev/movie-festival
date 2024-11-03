package utils

type CreateMovie struct {
	Title       string   `json:"title" validate:"required,max=50"`
	Description string   `json:"description"`
	Duration    int64    `json:"duration" validate:"required"`
	ArtistName  []string `json:"artist_name" validate"required"`
	GenreIds    []int16  `json:"genre_ids" validate:"required"`
	WatchUrl    string   `json:"watch_url" validate:"required"`
}
