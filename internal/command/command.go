package command

import (
	"fmt"
	"log"

	"github.com/Dhairya3124/blog-aggregator/internal/state"
)

type Command struct {
	Name string
	Args []string
}
type Commands struct{
	Handlers map[string]func(*state.State, Command) error
}

func handlerLogin(s *state.State,cmd Command)error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expects a single argument")
	}else{
		err:=s.Config.SetUser(cmd.Args[1])
		if err != nil {
			return err
		}else{
			fmt.Println("user has been set")
		}
	}

	return nil
}
func (c *Commands)register(name string,f func(*state.State,Command)error){
	c.Handlers[name] = f

}
func (c *Commands)run(s *state.State,cmd Command) error{
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
	return commands
}

