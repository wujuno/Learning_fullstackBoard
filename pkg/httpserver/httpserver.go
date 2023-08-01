package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartHTTPServer() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", getPostsHandler).Methods(http.MethodGet)
	
	/* r.HandleFunc("/signup", signupHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts:id", getPostHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/create", createPostHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts/delete", deletePostHandler).Methods(http.MethodDelete)
	r.HandleFunc("/comments/create", createCommentHandler).Methods(http.MethodPost) */
	
	


	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}