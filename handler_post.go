package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/natnael-alemayehu/grapton/internal/database"
)

func handlerBrowsePosts(s *state, cmd command, usr database.User) error {
	limit := 2
	if len(cmd.arg) != 0 {
		var err error
		limit, err = strconv.Atoi(cmd.arg[0])
		if err != nil {
			return err
		}
	}

	param := database.GetPostForUserParams{
		UserID: usr.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostForUser(context.Background(), param)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedsName)
		fmt.Printf("--- %s ---\n", post.PostTitle)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.PostUrl)
		fmt.Println("=====================================")
	}

	return nil
}
