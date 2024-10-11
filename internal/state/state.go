package state

import (
	"log"

	config "github.com/Dhairya3124/blog-aggregator/internal/config"
)

type State struct{
	Config *config.Config
}
func New() State {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	return State{
		Config: &cfg,
	}
}