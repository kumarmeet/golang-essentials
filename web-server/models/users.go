package models

import (
	"errors"
	"log"
	"time"

	"github.com/learning-webserver/db"
	"github.com/learning-webserver/utils"
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func (u *User) Save() (int64, error) {
	query := `
		INSERT INTO users(name, email, password) VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) GetUser(id int64) (*User, error) {
	query := "SELECT * FROM users where id = ?"

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Created_At, &u.Updated_At)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ValidCreds(email, password string) (*User, error) {
	query := "SELECT id, password FROM users where email = ?"

	row := db.DB.QueryRow(query, email)

	var retrivedPassword string
	var id int64

	err := row.Scan(&id, &retrivedPassword)

	if err != nil {
		return nil, errors.New("Credentials Invalid!")
	}

	paswordIsValid := utils.CheckPasswordHash(password, retrivedPassword)

	if !paswordIsValid {
		return nil, errors.New("Credentials Invalid!")
	}

	user, err := u.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
