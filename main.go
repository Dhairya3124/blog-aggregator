package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dhairya3124/blog-aggregator/internal/command"
	"github.com/Dhairya3124/blog-aggregator/internal/state"
)

func main() {
	s:=state.New()
	commands := command.NewCommands()
	args := os.Args
	fmt.Println(args)

	if len(args) < 2 {
		log.Fatal("Not enough arguments specified")
	}

	commands.Run(&s, command.Command{
		Name: args[1],
		Args: args[2:],
	})


	// configDetails, err := config.Read()
	// if err != nil {
	// 	log.Fatalf(" Following Error %v", err)
	// }
	// configDetails.SetUser("Dhairya")
	// configDetails, err = config.Read()
	// if err != nil {
	// 	log.Fatalf(" Following Error %v", err)
	// }
	fmt.Println(s.Config)

}
