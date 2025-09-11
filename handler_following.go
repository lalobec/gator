package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("A user must be logged in to follow a feed: %v\n", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return nil
	}
	for _, follow := range follows{
		fmt.Println(follow.FeedName)
	}
	return nil
}
