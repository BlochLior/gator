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
	fmt.Println("Successful reset of table.")
	return nil
}
