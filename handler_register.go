package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/BlochLior/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Fprintf(os.Stderr, "user %s already exists\n", name)
		os.Exit(1)
	}

	user, err := s.db.CreateUser(
		context.Background(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("failed setting current user %s in config", user.Name)
	}

	fmt.Println("User was created successfuly")
	printUser(user)
	return nil
}
