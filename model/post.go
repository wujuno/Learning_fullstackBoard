package model

type Post struct {
	PostId   int    `json:"id"`
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}