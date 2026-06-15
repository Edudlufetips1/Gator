package main

import (
    "log"
    "github.com/Edudlufetips1/Gator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config, %v", err)
	}
	err = cfg.SetUser("Asad")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config, %v", err)
	}
	fmt.Printf("%+v\n", cfg)
}