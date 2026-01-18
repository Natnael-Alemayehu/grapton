package main

import (
	"context"
	"fmt"
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

	for _, val := range rss.Channel.Item {
		fmt.Printf("  -%v \n", val.Title)
	}

	return nil
}
