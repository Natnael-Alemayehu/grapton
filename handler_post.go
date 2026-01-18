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

	for _, val := range posts {
		fmt.Printf("  - PostTitle: %v, \tpost link: %v, \tpublished at: %v\n", val.PostTitle, val.PostUrl, val.PublishedAt)
	}

	return nil
}
