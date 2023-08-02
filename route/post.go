package route

import (
	"encoding/json"
	"fmt"
	"fullstackboard/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["postId"])
	if err != nil {
		http.Error(w, "Invalid postId parameter", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	post, err := db.SelectPostInfo(postId)
	if err != nil {
		http.Error(w, "Failed to fetch post information", http.StatusInternalServerError )
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		if err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
}