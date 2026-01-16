package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natnael-alemayehu/grapton/internal/database"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arg) == 0 {
		return errors.New("login command expects an argument")
	}

	userName := cmd.arg[0]

	usr, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("User not found")
	}

	if err := s.cfg.SetUser(usr.Name); err != nil {
		return fmt.Errorf("Set User error: %v", err)
	}

	fmt.Println("user set successfully")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arg) == 0 {
		return errors.New("login command expects an argument")
	}

	userName := cmd.arg[0]

	NewUsr := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	usr, err := s.db.CreateUser(context.Background(), NewUsr)
	if err != nil {
		return err
	}

	s.cfg.CurrentUserName = userName
	s.cfg.SetUser(userName)

	fmt.Println("user registered successfully")
	fmt.Printf("New User Created: id: %v, created_at: %v, updated_at: %v, name: %v\n", usr.ID, usr.CreatedAt, usr.UpdatedAt, usr.Name)

	return nil
}

func handleListUsers(s *state, cmd command) error {
	usrs, err := s.db.QueryUsers(context.Background())
	if err != nil {
		return nil
	}

	currentUser := s.cfg.CurrentUserName

	for _, usr := range usrs {
		if usr.Name == currentUser {
			fmt.Printf("  * %v (current)\n", usr.Name)
			continue
		}
		fmt.Printf("  * %v\n", usr.Name)
	}

	return nil
}

func handleReset(s *state, _ command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("Reset Success!")
	return nil
}
