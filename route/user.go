package route

import (
	"encoding/json"
	"fmt"
	"fullstackboard/db"
	"fullstackboard/model"
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	

	err = db.InsertUser(&newUser)
	if err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created successfully")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	isExist, err := db.SelectExistUser(&user)
	if err != nil {
		http.Error(w, "Internal server error",  http.StatusInternalServerError)
		return
	}
	if isExist {
		w.WriteHeader(http.StatusOK)
        w.Write([]byte("로그인 성공"))
	} else {
		http.Error(w, "유저 정보가 일치하지 않습니다.", http.StatusUnauthorized)
	}
	
}