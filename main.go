package main

import (
    "log"
    "github.com/Edudlufetips1/Gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config, %v", err)
	}
	s := state {
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}
	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("not able to run command: %v", err)
	}
}