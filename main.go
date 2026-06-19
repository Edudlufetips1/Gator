package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Edudlufetips1/Gator/internal/config"
	"github.com/Edudlufetips1/Gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config, %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error retrieving database")
	}
	dbQueries := database.New(db)
	s := state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFeedFollow))

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("not able to run command: %v", err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
