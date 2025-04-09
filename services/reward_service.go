package services

import (
	"context"
	"errors"
	"pycore/config"
	"pycore/models"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ClaimReward(userID, code string) (*models.Reward, error) {
	ctx := context.Background()
	rewards := config.Client.Collection("rewards")

	query := rewards.Where("code", "==", code).Documents(ctx)
	doc, err := query.Next()
	if err == iterator.Done {
		return nil, errors.New("invalid reward code")
	} else if err != nil {
		return nil, err
	}

	var reward models.Reward
	doc.DataTo(&reward)

	if time.Now().After(reward.ExpiresAt) {
		return nil, errors.New("reward code has expired")
	}

	if reward.IsClaimed {
		return nil, errors.New("reward code has already been claimed")
	}

	_, err = doc.Ref.Update(ctx, []firestore.Update{
		{Path: "is_claimed", Value: true},
		{Path: "user_id", Value: userID},
	})
	if err != nil {
		return nil, err
	}

	reward.IsClaimed = true
	reward.UserID = userID

	return &reward, nil
}

func CreateReward(reward models.Reward) error {
	ctx := context.Background()
	rewards := config.Client.Collection("rewards")

	_, _, err := rewards.Add(ctx, map[string]interface{}{
		"id":          reward.ID,
		"code":        reward.Code,
		"description": reward.Description,
		"expires_at":  reward.ExpiresAt,
		"is_claimed":  false,
		"user_id":     "",
	})
	return err
}
