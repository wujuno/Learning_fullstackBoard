package db

import (
	"database/sql"
	"fullstackBoard/model"
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

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return posts, nil

}