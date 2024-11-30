package models

import (
	"errors"

	"github.com/oyen-bright/event_REST/db"
	"github.com/oyen-bright/event_REST/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {

	query := `
	INSERT INTO users(email, password)
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()

	return err

}

func (user *User) ValidateCredentials() error {

	query :=
		`
	SELECT  password, id FROM users where email = ? 
	`

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword, &user.ID)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
