package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Dhairya3124/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
func ScrapeFeeds(ctx context.Context, db *database.Queries, time_between_reqs string) {
duration,err:=time.ParseDuration(time_between_reqs)
if err != nil {
	fmt.Println("Invalid duration")
	return
}
fmt.Printf("Scraping feeds every %v",time_between_reqs)
ticker:=time.NewTicker(duration)
for;;<-ticker.C{
	feed,err:=db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println("Error in retrieving error")
		return
	}
	fetched,err:=FetchFeed(ctx,feed.Url)
	if err != nil {
		fmt.Println("Invalid url for feed")
		return
	}
	fmt.Println(fetched)
	for _,item:= range fetched.Channel.Item{
		_, err = db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			FeedID:      feed.ID,
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if err != nil {
			pgErr, ok := err.(*pq.Error)

				if !ok {
					log.Fatal("Error saving posts:", err.Error())
				}

				if ok && pgErr.Code.Name() != "unique_violation" {
					log.Fatal("Database error saving posts:", err.Error())

				}
		}
	}

}
}
