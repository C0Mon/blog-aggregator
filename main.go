package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/C0Mon/blog-aggregator/internal/config"
	"github.com/C0Mon/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg, err := config.Read()
	check(err)

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	// Initialise state
	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	// Initialise commands
	cmds := commands{
		commandMap: map[string]func(*state, command) error{},
	}

	// Register commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	// Get arguments
	userArgs := os.Args
	if len(userArgs) < 2 {
		log.Fatal(fmt.Errorf("Please enter a command"))
	}
	userCmd := command{
		name: userArgs[1],
	}
	if len(userArgs) < 3 {
		userCmd.arguments = []string{}
	} else {
		userCmd.arguments = userArgs[2:]
	}

	err = cmds.run(&s, userCmd)

	check(err)

}
