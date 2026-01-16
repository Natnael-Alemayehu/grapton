package main

import (
	"fmt"
	"log"

	"github.com/natnael-alemayehu/grapton/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error Reading: %v", err)
	}

	if err = cfg.SetUser("nate"); err != nil {
		log.Fatalf("Setting User Reading: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error Reading: %v", err)
	}

	fmt.Printf("DBURL: %v \nCurrent Username: %v\n", cfg.DBURL, cfg.CurrentUserName)
}
