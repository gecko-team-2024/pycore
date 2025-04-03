package models

import "time"

type OAuth struct {
	ID         string    `json:"id"`
	UserName   string    `json:"username"`
	Email      string    `json:"email"`
	PhotoURL   string    `json:"photo_url"`
	DayOfBirth time.Time `json:"day_of_birth"`
	CreatedAt  time.Time `json:"created_at"`
	PCoin      int       `json:"pcoin"`
	Role       string    `json:"role"`
	Method     string    `json:"method"`
}
