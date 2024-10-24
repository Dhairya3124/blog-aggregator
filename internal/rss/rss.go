package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	
		client := http.Client{}
	
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
		if err != nil {
			return nil, fmt.Errorf("error generating HTTP request: %v", err)
		}
	
		req.Header.Add("User-Agent", "gator")
	
		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error fetching RSS feed: %v", err)
		}
		defer res.Body.Close()
	
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}
	
		var feed RSSFeed
		err = xml.Unmarshal(data, &feed)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling XML feed: %v", err)
		}
	
		feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
		feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	
		for _, item := range feed.Channel.Item {
			item.Title = html.UnescapeString(item.Title)
			item.Description = html.UnescapeString(item.Description)
		}
	
		return &feed, nil
	
}
