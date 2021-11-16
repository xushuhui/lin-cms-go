package lib

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = "secret"

func GenerateJwt(claims jwt.MapClaims) (jwtToken string, err error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtToken, err = token.SignedString([]byte(jwtKey))
	return
}

func ParseJwt() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtKey),
	})
}
