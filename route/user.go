package route

import (
	"encoding/json"
	"fmt"
	"fullstackboard/db"
	"fullstackboard/model"
	"fullstackboard/util"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = string(hashedPassword)

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

	storedUser, err := db.SelectExistUser(&user)
	if err != nil {
		http.Error(w, "Internal server error",  http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "잘못된 비밀번호입니다.", http.StatusUnauthorized)
		return
	}

	accessToken, err := util.MakeTokenString(storedUser)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
    w.Write([]byte(accessToken))
}