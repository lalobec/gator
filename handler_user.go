package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lalobec/gator/internal/database"
	"time"
)

import _ "github.com/lib/pq"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}
	name := cmd.arguments[0]

	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("%s user does not exist", name)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("An error ocurred while login: %v\n", err)
	}

	fmt.Printf("%s is logged in the system \n", cmd.arguments[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	name := cmd.arguments[0]

	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, args.Name)
	if err == sql.ErrNoRows {
		user, errCreate := s.db.CreateUser(ctx, args)
		if errCreate != nil {
			return fmt.Errorf("Could not create user: %w\n", errCreate)
		}
		err2 := s.cfg.SetUser(args.Name)
		if err2 != nil {
			return fmt.Errorf("Could not set current user: %w\n", err2)
		}
		fmt.Println("User created successfully")
		printUser(user)
		return nil
	} else if err != nil {
		return err
	} else {
		return fmt.Errorf("username %v already exists \n", args.Name)
	}

}

func printUser(user database.User) {
	fmt.Printf(" * ID:			%v\n", user.ID)
	fmt.Printf(" * Name:			%v\n", user.Name)
}
