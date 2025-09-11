package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/lalobec/gator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	username, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("There is not a username logged in: %v\n", err)
	}

	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}

	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]
	args := database.GetFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:       feedName,
		Url:        feedURL,
		UserID: username.ID,
	}

	feed, err := s.db.GetFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("An error ocurred getting the feed: %v", err)
	}

	argsCreateFeed := database.CreateFeedFollowParams {
		UserID: username.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), argsCreateFeed)
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully")
	printFeed(feed)
	fmt.Println()
	fmt.Println("===============================")
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for idx, feed := range feeds {
		fmt.Println("=====================")
		fmt.Printf("Feed %d\n", idx)
		fmt.Println("=====================")
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		user, err := s.db.GetUserFromId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("UserID: %s\n", user.Name)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
