package command

import (
	"github.com/Dhairya3124/blog-aggregator/internal/state"
	"log"
)

type Command struct {
	Name string
	Args []string
}
type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f

}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler := c.Handlers[cmd.Name]

	err := handler(s, cmd)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
func NewCommands() Commands {
	commands := Commands{
		Handlers: make(map[string]func(*state.State, Command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregateRSSFeed)
	commands.register("addfeed", middlewareLoggedIn(handlerRSSFeed))
	commands.register("feeds", handlerShowRSSFeed)
	commands.register("follow", middlewareLoggedIn(handlerFollowRSSFeed))
	commands.register("following", middlewareLoggedIn(handlerFollowingRSSFeed))
	return commands
}
