package main

import (
	"fmt"
	"os"
)

func getArgs() (command, error) {
	if len(os.Args) < 2 {
		return command{}, fmt.Errorf("No arguments were given")
	}
	return command{
		name: os.Args[1],
		args: os.Args[2:],
	}, nil
}
