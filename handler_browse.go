package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AradD7/Gator/internal/database"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("The command browse takes 1 optional argument: limit")
	}
	var browseLimit int32 = 2
	if len(cmd.args) == 1 {
		num64, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("The browse limit must a number")
		}
		browseLimit = int32(num64)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: browseLimit,
	})
	if err != nil {
		return fmt.Errorf("Failed getting posts: %v", err)
	}

	if len(posts) == 0 {
		return fmt.Errorf("No post is available")
	}

	for _, post := range posts {
		fmt.Printf("Title:   %v\n", post.Title.String)
		fmt.Printf("url:     %v\n\n", post.Url)
	}
	return nil
}
