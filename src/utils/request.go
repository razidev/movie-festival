package utils

type PayloadMovie struct {
	Title       string   `json:"title" validate:"required,max=50"`
	Description string   `json:"description"`
	Duration    int64    `json:"duration" validate:"required"`
	ArtistName  []string `json:"artist_name" validate"required"`
	GenreIds    []int16  `json:"genre_ids" validate:"required"`
	WatchUrl    string   `json:"watch_url" validate:"required,http_url"`
}

type PayloadUser struct {
	Email    string `json:"email" validate:"required,min=5,max=20,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
