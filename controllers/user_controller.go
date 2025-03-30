package controllers

import (
	"encoding/json"
	"net/http"
	"pycore/models"
	"pycore/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	newUser, err := services.RegisterWithEmailAndPassword(
		user.Email,
		user.Password,
		user.UserName,
	)

	if err != nil {
		http.Error(w, "Can't create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	userID, err := services.LoginWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"userID": userID, "message": "Login successfully"})
}
