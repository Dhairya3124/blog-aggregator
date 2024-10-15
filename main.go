package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Dhairya3124/blog-aggregator/internal/command"
	"github.com/Dhairya3124/blog-aggregator/internal/database"
	"github.com/Dhairya3124/blog-aggregator/internal/state"
	_ "github.com/lib/pq"
)

func main() {
	s := state.New()
	db, err := sql.Open("postgres", s.Config.DbURL)
	if err != nil {
		log.Fatal("Error in connecting database!")
	}
	defer db.Close()
	dbQueries := database.New(db)
	s.DB = dbQueries
	commands := command.NewCommands()
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments specified")
	}
	commands.Run(&s, command.Command{
		Name: args[1],
		Args: args[2:],
	})
	

}
