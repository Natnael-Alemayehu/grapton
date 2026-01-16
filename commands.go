package main

import (
	"fmt"
)

type command struct {
	name string
	arg  []string
}

type commands struct {
	commands map[string]func(s *state, cmd command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}

	return handler(s, cmd)
}
