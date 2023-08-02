package route

import (
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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	newUser := model.User{
		Name: username,
		Email: email,
		Password: password,
	}

	err = db.InsertUser(&newUser)
	if err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusCreated)


}