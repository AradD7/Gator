package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AradD7/Gator/internal/database"
	"github.com/AradD7/Gator/internal/fetch"
	"github.com/google/uuid"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("The agg command expects 1 argument: time_between_reqs")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("The agg command expects a time: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}


func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Something went wrong getting the next feed to fetch: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("Something went wrong marking the feed at fetched: %v", err)
	}

	rssFeed, err := fetch.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Something went wrong fetching the feed: %v", err)
	}

	for _, rssItem := range rssFeed.Channel.Item {
		if _, err := s.db.GetPostByURL(context.Background(), rssItem.Link); err == nil {
			fmt.Println("found an already added post")
			continue
		}
		itemTitle := sql.NullString {
			String: rssItem.Title,
			Valid:  rssItem.Title != "",
		}
		itemDescription := sql.NullString {
			String: rssItem.Description,
			Valid:  rssItem.Description != "",
		}
		itemPubDate, err := time.Parse(time.RFC1123Z, rssItem.PubDate)
		if err != nil {
			itemPubDate, _ = time.Parse(time.RFC1123, rssItem.PubDate)
		}

		post := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: itemTitle,
			Url: rssItem.Link,
			Description: itemDescription,
			PublishedAt: itemPubDate,
			FeedID: feed.ID,
		}

		newPost, err := s.db.CreatePost(context.Background(), post)
		if err != nil {
			return err
		}
		fmt.Printf("\nNew post has been added!\n")
		fmt.Printf("ID:           %v\n", newPost.ID)
		fmt.Printf("Title:        %v\n", newPost.Title.String)
		fmt.Printf("URL:          %v\n", newPost.Url)
		fmt.Printf("Description:  %v\n\n", newPost.Description.String)
	}

	return nil
}
