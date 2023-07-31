package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartHTTPServer() {
	r := mux.NewRouter()

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}