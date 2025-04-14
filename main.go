package main

import (
	"fmt"
	"log"

	"github.com/BlochLior/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error trying to read config: %s", err)
	}

	err = cfg.SetUser("BlochLior")
	if err != nil {
		log.Fatalf("error trying to set user: %s", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error trying to re-read config: %s", err)
	}
	fmt.Print(cfg)
}
