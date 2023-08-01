package httpserver

import (
	"encoding/json"
	"fmt"
	"fullstackBoard/db"
	"net/http"
)

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")

	//db에서 가져오기
	posts, err := db.SelectPostsInfo()
	if err != nil {
		http.Error(w, "Failed to fetch posts information", http.StatusInternalServerError )
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		if err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
}