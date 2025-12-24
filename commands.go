package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/C0Mon/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdToRun, found := c.commandMap[cmd.name]
	if !found {
		return fmt.Errorf("Command Not Found")
	}
	return cmdToRun(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	err = nil

	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err == nil {
		fmt.Printf("user already exists in database\n")
		os.Exit(1)
	}
	err = nil

	currentTime := time.Now()
	userData, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
			Name:      cmd.arguments[0],
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", userData)
	return nil
}
