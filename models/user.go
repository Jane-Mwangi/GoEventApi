package models

import (
	"errors"

	"github.com/Jane-Mwangi/GoEventApi/db"
	"github.com/Jane-Mwangi/GoEventApi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required" `
}

func (u User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
    query := "SELECT id, password FROM users WHERE email = ?"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    err := row.Scan(&u.ID, &retrievedPassword)
    if err != nil {
        return errors.New("Invalid credentials")
    }

    fmt.Println("Provided password:", u.Password)
    fmt.Println("Retrieved hashed password:", retrievedPassword)

    passwordIsValid := utils.ComparePassword(u.Password, retrievedPassword)
    fmt.Println("Password is valid:", passwordIsValid)

    if !passwordIsValid {
        return errors.New("Invalid credentials")
    }

    return nil
}
