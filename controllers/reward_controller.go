package controllers

import (
	"encoding/json"
	"net/http"
	"pycore/services"
)

type ClaimRewardRequest struct {
	UserID string `json:"user_id"`
	Code   string `json:"code"`
}

func ClaimRewardHandler(w http.ResponseWriter, r *http.Request) {
	var req ClaimRewardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reward, err := services.ClaimReward(req.UserID, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reward)
}
