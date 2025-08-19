package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AradD7/Gator/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 2 {
		return fmt.Errorf("The add feed command expects 2 arguments: name and url")
	}


	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:  		cmd.args[0],
		Url:  		cmd.args[1],
		UserID:  	user.ID,
	})
	if err != nil {
		return fmt.Errorf("Something went wrong adding the feed to the database, %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: newFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("Something went wrong following the feed, %v", err)
	}

	fmt.Println("Feed successfully created and followed!")
	printFeed(newFeed)
	return nil
}


func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("The feeds command expects no arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Something went wrong getting the feeds: %v", err)
	}
	if len(feeds) == 0 {
		fmt.Print("There are no feeds stored")
		return nil
	}

	fmt.Println("---------------------------------------")
	for _, feed := range feeds {
		createdBy, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			createdBy.Name = "user not found"
		}
		fmt.Printf("Name:        %v\n", feed.Name)
		fmt.Printf("URL:         %v\n", feed.Url)
		fmt.Printf("Created By:  %v\n", createdBy.Name)
		fmt.Println("---------------------------------------")
	}
	return nil
}


func printFeed(feed database.Feed) {
	fmt.Printf("ID:          %v\n", feed.ID)
	fmt.Printf("Created_At:  %v\n", feed.CreatedAt)
	fmt.Printf("Updated_At:  %v\n", feed.UpdatedAt)
	fmt.Printf("Name:        %v\n", feed.Name)
	fmt.Printf("URL:         %v\n", feed.Url)
	fmt.Printf("User_ID:     %v\n", feed.UserID)
}
