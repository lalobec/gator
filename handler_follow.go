package main

import (
	"context"
	"fmt"
	"github.com/lalobec/gator/internal/database"
	"github.com/google/uuid"
	"time"	
)

import _ "github.com/lib/pq"

func handlerFollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, args)
	if err != nil {
		return err
	}
	fmt.Println("Following!")
	fmt.Printf("user %v", user.Name)
	fmt.Printf("feed %v", feed.Name)
	return nil
}
