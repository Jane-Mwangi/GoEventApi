package utils

import (
	"errors"
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

func VerifyToken(token string)(int64, error ){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(secretKey), nil

	})

	if err != nil {
		return 0, errors.New("could not parse token")

	}
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0,errors.New("Token is not valid")
	}

	claims,ok:=parsedToken.Claims.(jwt.MapClaims)

	if !ok{
		return 0, errors.New("Invalid	id token claims")
	}

	// email:=claims["email"].(string)
	userIdFloat, ok := claims["userId"].(float64)
    if !ok {
        return 0, errors.New("userId is not a valid float64")
    }

    userId := int64(userIdFloat)
	return  userId,nil

}
