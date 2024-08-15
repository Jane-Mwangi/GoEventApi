package utils

import (
	"errors"
	"os/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email, // Use string literals for the keys
		//never include sensitive information in the token e.g passwords
		"userId": userId,
		//expriration time
		"exp": time.Now().Add(time.Hour * 2).Unix(), // Expiration time
	})

	return token.SignedString([]byte(secretKey)) // Use the secretKey constant
}

func VerifyToken(token string) error {
	parsedToken,err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil, errors.New("Unexpected signing method: %v", token.Header["alg"])
		}
		return (secretKey), nil

	})

	if err!=nil{
		return errors.New("could not parse token")
		
	}
	tokenIsValid:=parsedToken.Valid

	if !tokenIsValid{
		return errors.New("Token is not valid")
	}

	// claims,ok:=parsedToken.Claims(jwt.MapClaims)

	// if !ok{
	// 	return errors.New("Invalid token claims")
	// }

	// email:=claims["email"].(string)
	// userId:=claims["userId"].(int64)
	return nil

}
