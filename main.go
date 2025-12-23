package main

import (
	"fmt"
	"log"
	"os"

	"github.com/C0Mon/blog-aggregator/internal/config"
)

type state struct {
	Config *config.Config
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg, err := config.Read()
	check(err)

	// Initialise state
	s := state{
		Config: &cfg,
	}

	// Initialise commands
	cmds := commands{
		commandMap: map[string]func(*state, command) error{},
	}

	// Register commands
	cmds.register("login", handlerLogin)

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
