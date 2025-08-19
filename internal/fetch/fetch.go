package fetch

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
)



type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func (r RSSFeed) cleanHTML() {
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (r RSSItem) cleanHTML() {
	r.Description = html.UnescapeString(r.Description)
	r.Title = html.UnescapeString(r.Title)
}


func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, strings.NewReader(""))
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to create GET request: %v", err)
	}
	req.Header.Set("User-Agen", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to send GET request: %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to read the data returned by server: %v", err)
	}

	var feed RSSFeed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to unmarshal the XML data: %v", err)
	}

	feed.cleanHTML()
	for _, item := range feed.Channel.Item {
		item.cleanHTML()
	}

	return &feed, nil
}


