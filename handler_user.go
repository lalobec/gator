package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("You must provide a username")
	}
	name := cmd.arguments[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("An error ocurred while login: %v\n", err)
	}
	
	fmt.Printf("%s is logged in the system \n", cmd.arguments[0])
	return nil
}
