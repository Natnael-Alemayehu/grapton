package main

import (
	"context"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	updatedFeed, err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rss, err := fetchFeed(context.Background(), updatedFeed.Url)
	if err != nil {
		return err
	}

	for _, post := range rss.Channel.Item {
		pAt, err := dateparse.ParseAny(post.PubDate)
		if err != nil {
			return fmt.Errorf("post published at formatting error: %v", err)
		}
		postParam := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: pAt,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), postParam)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			return err
		}

	}
	return nil
}
