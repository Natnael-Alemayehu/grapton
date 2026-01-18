package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

func handlerCreateFeedFolowers(s *state, cmd command, usr database.User) error {

	if len(cmd.arg) == 0 {
		return errors.New("feed to follow not provided")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arg[0])
	if err != nil {
		return err
	}

	cff := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feed.ID,
	}

	createdFeedFollowers, err := s.db.CreateFeedFollows(context.Background(), cff)
	if err != nil {
		return err
	}

	fmt.Printf("Created Feed: %v\n", createdFeedFollowers)
	return nil
}

func handlerGetFollowersForUser(s *state, _ command, usr database.User) error {

	followForUser, err := s.db.GetFeedFollowsForUser(context.Background(), usr.ID)
	if err != nil {
		return err
	}

	for _, v := range followForUser {
		fmt.Printf("  - %v\n", v.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, usr database.User) error {
	if len(cmd.arg) == 0 {
		return errors.New("unfollow needs a feed url to unfollow")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arg[0])
	if err != nil {
		return err
	}

	unfParam := database.UnfollowFeedParams{
		FeedID: feed.ID,
		UserID: usr.ID,
	}

	return s.db.UnfollowFeed(context.Background(), unfParam)
}
