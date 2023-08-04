package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartHTTPServer() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", getPostsHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/{postId}", getPostHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts", createPostHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts/{postId}", deletePostHandler).Methods(http.MethodDelete)
	r.HandleFunc("/posts/{postId}", updatePostHandler).Methods(http.MethodPost)
	
	r.HandleFunc("/users", signupHandler).Methods(http.MethodPost)
/* 	r.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	r.HandleFunc("/comments/create", createCommentHandler).Methods(http.MethodPost) */
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}