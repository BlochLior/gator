package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("failure resetting table: %s", err)
		return err
	}
	fmt.Println("Successful reset of users table.")
	err = s.db.ResetFeeds(context.Background())
	if err != nil {
		fmt.Printf("failure resetting feeds: %s", err)
		return err
	}
	fmt.Println("Successful reset of feeds table.")
	err = s.db.ResetFeedFollows(context.Background())
	if err != nil {
		fmt.Printf("failure resetting feed_follows: %s", err)
		return err
	}
	fmt.Println("Successful reset of feed follows table.")

	return nil
}
