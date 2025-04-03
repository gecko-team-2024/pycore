package services

import (
	"context"
	"errors"
	"fmt"
	"pycore/config"
	"pycore/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Error: ", err
	}
	return string(hashes), nil
}

func RegisterWithEmailAndPassword(email, password, username string) (string, error) {
	ctx := context.Background()

	if email == "" || password == "" || username == "" {
		return "", errors.New("email, password, and username are required")
	}

	users := config.Client.Collection("users")
	query := users.Where("email", "==", email).Documents(ctx)
	existingUser, _ := query.Next()
	if existingUser != nil {
		return "", errors.New("email already exists")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return "Error hashed password", err
	}

	userID := uuid.New().String()

	newUser := models.User{
		ID:        userID,
		UserName:  username,
		Password:  hashedPassword,
		Email:     email,
		CreatedAt: time.Now(),
		PCoin:     0,
		Role:      "user",
		Method:    "email/password",
		PhotoURL:  fmt.Sprintf("https://ui-avatars.com/api/?name=%s&background=random&color=fff", username),
	}

	_, err = users.Doc(userID).Set(ctx, newUser)
	if err != nil {
		return "Error when create new user: ", err
	}

	return userID, nil
}

func LoginWithEmailAndPassword(email, password string) (string, error) {
	ctx := context.Background()

	if email == "" || password == "" {
		return "", errors.New("email and password cannot be blank")
	}

	users := config.Client.Collection("users")
	query := users.Where("Email", "==", email).Documents(ctx)
	doc, err := query.Next()

	if err != nil {
		return "", errors.New("email does not exist or password is incorrect")
	}

	var user struct {
		ID       string `firestore:"id"`
		Password string `firestore:"password"`
	}

	doc.DataTo(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("password is not correct")
	}

	return user.ID, nil
}

func GetUserByID(userId string) (*models.User, error) {
	doc, err := config.Client.Collection("users").Doc(userId).Get(context.Background())
	if err != nil {
		return nil, errors.New("user not found")
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, errors.New("false to parse user data")
	}

	return &user, nil
}
