package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey="supersecret"

func GenerateToken(email string, userId string) (string, error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		email: email,
		//never include sensitive information in the token e.g password
		userId: userId,
		//exp is the time the token will expire
		exp: time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte("secretKey"))
}
