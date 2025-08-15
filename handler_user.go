package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AradD7/Gator/internal/database"
	"github.com/google/uuid"
)


func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Register command expects a single name")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		return fmt.Errorf("User %s already exists!", cmd.args[0])
	}
	createdUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: 	uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	})
	if err != nil {
		return err
	}
	if err = s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("User was created successfully!")
	printUser(createdUser)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("The login command expects 1 username")
	}
	username := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), username); err != nil {
		return fmt.Errorf("Couldn't find username in the database")
	}
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("The user has been set!")
	return nil
}


func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}


func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("the reset command expects no arguments")
	}
	if err := s.db.Reset(context.Background()); err != nil {
		return err
	}
	log.Printf("database was successfully reset")
	return nil
}


func handlerLogUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("the users command expects no arguments")
	}
	currentUser := s.cfg.CurrentUserName
	Users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range Users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
