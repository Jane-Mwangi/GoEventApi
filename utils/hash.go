package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error) {

	bytse,err :=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytse),err
}

func ComparePassword(password , hashedPassword string) bool {
	err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))

	return err == nil 
}