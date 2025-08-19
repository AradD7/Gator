package main

import (
	"database/sql"
	"log"

	"github.com/AradD7/Gator/internal/config"
	"github.com/AradD7/Gator/internal/database"
	_ "github.com/lib/pq"
)


func main() {
	ptr, err := config.Read()
	if err != nil {
		log.Fatalf("error reading the config: %v", err)
	}

	db, err := sql.Open("postgres", ptr.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	dbQueries := database.New(db)

	cfgState := state{
		cfg: &ptr,
		db: dbQueries,
	}

	commandsMap := commands{
		commandMap: map[string]func(*state, command) error{},
	}

	commandsMap.register("login", handlerLogin)
	commandsMap.register("register", handlerRegister)
	commandsMap.register("reset", handlerReset)
	commandsMap.register("users", handlerLogUsers)
	commandsMap.register("agg", handlerAgg)
	commandsMap.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commandsMap.register("feeds", handlerFeeds)
	commandsMap.register("follow", middlewareLoggedIn(handlerFollow))
	commandsMap.register("following", middlewareLoggedIn(handlerFollowing))
	commandsMap.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commandsMap.register("browse", middlewareLoggedIn(handlerBrowse))

	stdinArgs, err := getArgs()
	if err != nil {
		log.Fatalf("No command arguments were given")
	}

	if err = commandsMap.run(&cfgState, stdinArgs); err != nil {
		log.Fatalf("error running the command: %v", err)
	}
}
