package main

import (
	"context"
	"fmt"
	"github.com/lalobec/gator/internal/database"
)

import _ "github.com/lib/pq"

func handlerFollow(s *state, cmd command) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("A user must be logged in to follow a feed: %v\n", err)
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
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
