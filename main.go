package main

import (
    "log"
    "github.com/Edudlufetips1/Gator/internal/config"
	"github.com/Edudlufetips1/Gator/internal/database"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
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
	s := state {
		db: dbQueries,
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
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFeedFollow)

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("not able to run command: %v", err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	}
