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

func SelectExistUser(user *model.User) (*model.User, error) {
	q := `
		SELECT user_id, username, password
		FROM users
		WHERE username = ?;
	`
	var storedUser model.User
	err := DB.QueryRow(q, user.Name).Scan(&storedUser.Id, &storedUser.Name, &storedUser.Password)
	if err != nil {
		return nil, err
	}
	
	return &storedUser, nil
}