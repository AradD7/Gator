package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AradD7/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("The follow command expects 1 argument: url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Could not find a feed with the given url: %v", err)
	}


	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Something went wrong creating the feedfollow: %v", err)
	}

	fmt.Println("Feed has been successfully followed!")
	printFeedFollow(feedFollow)

	return nil
}


func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("The following command expects no argument")
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Something went wrong getting the followed feeds: %v", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("You are not following any feeds. To follow a feed use follow {feed_URL} command")
	} else{
		for _, feed := range feedFollows {
			fmt.Printf("* %s\n", feed.FeedName)
		}
	}

	return nil
}


func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("The unfollow command expects 1 argument: url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Could not find a feed with the given url: %v", err)
	}

	if err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("Failed to unfollow the feed: %v", err)
	}

	fmt.Println("Feed successfully unfollowed!")
	return nil
}


func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("ID:           %v\n", feedFollow.ID)
	fmt.Printf("Created_at:   %v\n", feedFollow.CreatedAt)
	fmt.Printf("updated_at:   %v\n", feedFollow.UpdatedAt)
	fmt.Printf("Followed by:  %v\n", feedFollow.UserName)
	fmt.Printf("Feed Name:    %v\n", feedFollow.FeedName)
}


