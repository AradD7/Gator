package main

import (
	"github.com/AradD7/Gator/internal/config"
	"github.com/AradD7/Gator/internal/database"
)

type state struct {
	db 	*database.Queries
	cfg *config.Config
}


