package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

func handerFeedAggregator(s *state, cmd command) error {
	if len(cmd.arg) == 0 {
		return errors.New("agg expects a time between requests like 1s, 5s, 1m, 5m ...")
	}

	timereq, err := time.ParseDuration(cmd.arg[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v \n", timereq)

	ticker := time.NewTicker(timereq)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func handlerAddFeed(s *state, cmd command, usr database.User) error {
	if len(cmd.arg) != 2 {
		return fmt.Errorf("addfeed argument needs name and feedurl")
	}

	name := cmd.arg[0]
	url := cmd.arg[1]

	dbfeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    usr.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), dbfeed)
	if err != nil {
		return err
	}

	cff := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    newFeed.ID,
	}
	createdFeed, err := s.db.CreateFeedFollows(context.Background(), cff)
	if err != nil {
		return fmt.Errorf("Create Feed Follows: %v", err)
	}

	fmt.Println("Feed successfully added")
	fmt.Printf("Feed successfully added to add feed: %v", createdFeed.FeedName)
	fmt.Printf("%v\n", newFeed)

	return nil
}

func handlerListFeeds(s *state, _ command) error {
	feeds, err := s.db.FeedDetail(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("FeedName: %v\n", feed.FeedName)
		fmt.Printf("\t- URL: %v    - name: %v\n", feed.FeedUrl, feed.UserName)
	}

	return nil
}
