package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/natnael-alemayehu/grapton/internal/config"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error Reading: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("postgrs connection error: %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		commands: make(map[string]func(s *state, cmd command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleListUsers)
	cmds.register("agg", handerFeedAggregator)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerCreateFeedFolowers))
	cmds.register("following", middlewareLoggedIn(handlerGetFollowersForUser))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name: os.Args[1],
		arg:  os.Args[2:],
	}

	if err := cmds.run(programState, cmd); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
