package db

import "fullstackBoard/model"

func InsertUser(user *model.User) error {
	q :=`
		INSERT INTO users(user_id, username, user_email, password)
		VALUES (Null, ?, ?, ?)
	`
	_, err := DB.Exec(q, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}