package main

import "fmt"

type command struct {
	name 	string
	args	[]string
}

type commands struct {
	commandMap 	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.commandMap[cmd.name]; !ok {
		return fmt.Errorf("Command not found")
	} else {
		return handler(s, cmd)
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
