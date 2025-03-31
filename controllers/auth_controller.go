package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"pycore/config"
	"pycore/models"
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOAuthConfig.AuthCodeURL("random-state")
	http.Redirect(w, r, url, http.StatusFound)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No code in query", http.StatusBadRequest)
		return
	}

	token, err := config.GoogleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	googleUser, err := fetchGoogleUser(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to fetch user info", http.StatusInternalServerError)
		return
	}

	googleUser.Method = "google"

	ctx := context.Background()
	user, err := getOrCreateUser(ctx, googleUser)
	if err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func fetchGoogleUser(accessToken string) (*models.OAuth, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var googleResponse struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		DisplayName   string `json:"name"`
		PhotoURL      string `json:"picture"`
		VerifiedEmail bool   `json:"verified_email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleResponse); err != nil {
		return nil, err
	}

	// Chuyển đổi dữ liệu thành `models.OAuth`
	user := &models.OAuth{
		ID:       googleResponse.ID,
		UserName: googleResponse.DisplayName,
		Email:    googleResponse.Email,
		PhotoURL: googleResponse.PhotoURL,
		Method:   "google",
		Role:     "user",
	}

	return user, nil
}

func getOrCreateUser(ctx context.Context, googleUser *models.OAuth) (*models.OAuth, error) {
	userRef := config.Client.Collection("users").Doc(googleUser.ID)
	doc, _ := userRef.Get(ctx)

	if !doc.Exists() {
		_, err := userRef.Set(ctx, googleUser)
		if err != nil {
			return nil, err
		}
	}

	return googleUser, nil
}
