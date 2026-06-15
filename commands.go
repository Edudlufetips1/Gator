package main

import (
	"errors"
	"fmt"
	"github.com/Edudlufetips1/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires at least one argument")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %s has been set!\n", cmd.args[0])
	return nil
}
