package main

import (
	"context"
	"fmt"
	"github.com/lalobec/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	follows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return nil
	}
	for _, follow := range follows{
		fmt.Println(follow.FeedName)
	}
	return nil
}
