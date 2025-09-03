package main

import (
	"fmt"
	"log"
	"github.com/lalobec/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	
	if err := cfg.SetUser("lalobg"); err != nil {
		log.Fatalf("Could not set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config once more: %+v\n", cfg)
}
