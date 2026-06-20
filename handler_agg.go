package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Edudlufetips1/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("at least one argument required")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("unable to fetch feed")
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking %s as fetched\n", feed.ID)
		return
	}
	fmt.Printf("Fetched feed %s has been marked\n", feed.ID)
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("FETCH ERROR: %v\n", err)
		return
	}
	fmt.Println("CHECKPOINT: about to loop")
	for _, item := range rssFeed.Channel.Item {
		log.Printf("starting loop iteration")
		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			parsedTime, err = time.Parse(time.RFC3339, item.PubDate)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: parsedTime, Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Printf("error creating post: %v\n", err)
			continue
		}
		log.Printf("saved post: %s\n", item.Title)
	}
}
