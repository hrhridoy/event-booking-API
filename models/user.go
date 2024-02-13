package models

import (
	"errors"

	"github.com/hrhridoy/event-booking-API/db"
	"github.com/hrhridoy/event-booking-API/utils"
)

type User struct {
	ID       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

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
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userID

	return nil

}

func (u *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CompareHashedPass(retrievedPassword, u.Password)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
