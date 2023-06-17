package Models

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID uint `json:"userID"`
	jwt.StandardClaims
}

func GenerateJWTToken(userID uint) (string, error) {
	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiry time: 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("asdffdasasdffasd")) // Replace "secret" with your actual secret key
}
