package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlochLior/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	err := printFeeds(s)
	if err != nil {
		return err
	}
	return nil
}
func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	feedURL := cmd.Args[1]
	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return err
	}
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       feedURL,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfuly:")
	printFeed(newFeed)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
func printFeeds(s *state) error {
	feeds, err := s.db.GetFeedsSpecial(context.Background())
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		fmt.Printf("Feed %d\n", i)
		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)
		fmt.Printf("* UserName:      %v\n", feed.UserName.String)
		fmt.Println()
	}
	return nil
}

//
//
//
// list feeds ready to test when i come back
