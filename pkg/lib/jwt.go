package lib

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("setting.JwtSecret")

type Claims struct {
	Uid uint `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid uint) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "skrshop",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(authString string) (*Claims, error) {

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil, errors.New("token error")
	}
	token := kv[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
