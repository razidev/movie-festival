package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Email    string    `json:"username"`
	UniqueID uuid.UUID `json:"unique_id"`
	jwt.StandardClaims
}

func GenerateJWT(email string, uniqueId uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email:    email,
		UniqueID: uniqueId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
