package main

import (
	"context"
	"fmt"
	"github.com/lalobec/gator/internal/database"
)

func middlewareLoggedIn(handler func(*state, command, database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("No user is logged in")
		}
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("A user must be logged in to follow a feed: %v\n", err)
		}
		return handler(s, cmd, user)
	}
}
