package state

import (
	"log"

	config "github.com/Dhairya3124/blog-aggregator/internal/config"
	"github.com/Dhairya3124/blog-aggregator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
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
