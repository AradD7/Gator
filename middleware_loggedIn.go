package main

import (
	"context"
	"fmt"

	"github.com/AradD7/Gator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Couldn't get current user from database: %v", err)
		}
		return handler(s, cmd, currentUser)
	}
}
