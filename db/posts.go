package db

import (
	"database/sql"
	"errors"
	"fmt"
	"fullstackboard/model"
)

func SelectPostsInfo() ([]model.Post, error) {
	posts := []model.Post{}

	q := `
		SELECT *
		FROM posts
	`

	rows, err := DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		post := model.Post{}
		var content sql.NullString
		
		err := rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.Title,
			&content,
		)

		if err != nil {
			return nil, err
		}

		if content.Valid {
			post.Content = content.String
		} else {
			post.Content = ""
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func SelectPostInfo(postId int) (model.Post, error) {
	q := `
		SELECT *
		FROM posts
		WHERE post_id = ?
	`
	row := DB.QueryRow(q, postId)
	
	post := model.Post{}
	var content sql.NullString

	err := row.Scan(
		&post.PostId,
		&post.UserId,
		&post.Title,
		&content,
	)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func InsertPost(post *model.Post) error {
	//TODO: else문 지워도 됨.
	if post.Title == "" || post.Content == "" {
		return errors.New("Title and content cannot be empty")
	} else{
		fmt.Println("title:",post.Title)
		fmt.Println("content:",post.Content)
	}

	//TODO: user_id 파라미터 값 수정해야함. 
	q := `
		INSERT INTO posts(post_id, user_id, title, content)
		VALUES (NULL, 1, ?, ?)
	`
	_, err := DB.Exec(q, post.Title, post.Content)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id int) error {
	q := `
		DELETE FROM posts
		WHERE post_id = ?
	`

	_, err := DB.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}