package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arg) == 0 {
		return errors.New("login command expects an argument")
	}

	userName := cmd.arg[0]

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("Set User error: %v", err)
	}

	fmt.Println("user set successfully")

	return nil
}
