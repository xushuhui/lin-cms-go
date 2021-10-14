package lib

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(claims jwt.MapClaims) (jwtToken string, err error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtToken, err = token.SignedString([]byte("secret"))
	return
}
