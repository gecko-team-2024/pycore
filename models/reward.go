package models

import "time"

type Reward struct {
	ID          string    `json:"id" firestore:"id"`
	Code        string    `json:"code" firestore:"code"`
	Description string    `json:"description" firestore:"description"`
	ExpiresAt   time.Time `json:"expires_at" firestore:"expires_at"`
	IsClaimed   bool      `json:"is_claimed" firestore:"is_claimed"`
	UserID      string    `json:"user_id" firestore:"user_id"`
}
