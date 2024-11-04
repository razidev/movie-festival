package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func ParseJWT(c *gin.Context) (*Claims, bool) {
	claims := &Claims{}
	tokenString := c.GetHeader("Authorization")

	// Remove the "Bearer " prefix if it exists
	tokenString = removeBearerPrefix(tokenString)

	// Parse the JWT with the claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			return nil, false
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
		return nil, false
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil, false
	}

	// Check if the token has expired
	if claims.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return nil, false
	}

	return claims, true
}

// Helper function to remove "Bearer " prefix from the token string if it's present
func removeBearerPrefix(tokenString string) string {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		return tokenString[7:]
	}
	return tokenString
}
