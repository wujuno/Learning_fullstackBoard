package model

type Post struct {
	PostId   int    `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}