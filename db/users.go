package db

import (
	"errors"
	"fmt"
	"fullstackboard/model"
)

func InsertUser(user *model.User) error {
	if user.Name == "" || user.Password == "" {
		return errors.New("Name and Password cannot be empty")
	} else{
		fmt.Println("Name:",user.Name)
		fmt.Println("Password:",user.Password)
	}

	q := `
		INSERT INTO users(user_id, username, user_email, password)
		VALUES (Null, ?, ?, ?)
	`
	_, err := DB.Exec(q, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func SelectExistUser(user *model.User) (string, error) {
	q := `
		SELECT password
		FROM users
		WHERE username = ?;
	`
	var hashedPassword string
	err := DB.QueryRow(q, user.Name).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}
	
	return hashedPassword, nil
}