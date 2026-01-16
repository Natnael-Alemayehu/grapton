package main

import (
	"log"
	"os"

	"github.com/natnael-alemayehu/grapton/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error Reading: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		commands: make(map[string]func(s *state, cmd command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name: os.Args[1],
		arg:  os.Args[2:],
	}

	if err := cmds.run(s, cmd); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
