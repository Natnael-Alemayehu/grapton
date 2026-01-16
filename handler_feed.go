package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

func handerFeedAggregator(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arg) != 2 {
		return fmt.Errorf("addfeed argument needs name and feedurl")
	}

	name := cmd.arg[0]
	url := cmd.arg[1]

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("GetUser in addfeed error: %v", err)
	}

	dbfeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    usr.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), dbfeed)
	if err != nil {
		return err
	}

	fmt.Println("Feed successfully added")
	fmt.Printf("%v\n", feed)

	return nil
}
