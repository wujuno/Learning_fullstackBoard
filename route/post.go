package route

import (
	"encoding/json"
	"fmt"
	"fullstackboard/db"
	"fullstackboard/model"
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

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var newPost model.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	err = db.InsertPost(&newPost)
	if err != nil {
		http.Error(w, "Failed to create new post", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Post created successfully")
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["postId"])
	if err != nil {
		http.Error(w, "Invalid postId parameter", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	
	err = db.DeletePost(postId)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError )
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Post deleted successfully")
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var UpdatedPost model.Post
	err := json.NewDecoder(r.Body).Decode(&UpdatedPost)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["postId"])
	if err != nil {
		http.Error(w, "Invalid postId parameter", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	err = db.UpdatePost(&UpdatedPost, postId)
	if err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Post updated successfully")
}