package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dhairya3124/blog-aggregator/internal/database"
	"github.com/Dhairya3124/blog-aggregator/internal/state"
	"github.com/Dhairya3124/blog-aggregator/internal/rss"
	"github.com/google/uuid"
)

type Command struct {
	Name string
	Args []string
}
type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func handlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expects a single argument")
	} else {
		username := cmd.Args[0]
		user, err := s.DB.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("error fetching user: %v", err)
		}

		err = s.Config.SetUser(user.Name)
		if err != nil {
			return err
		} else {
			fmt.Println("user has been set")
		}

	}

	return nil
}
func handlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expects a single argument")
	} else {
		username := cmd.Args[0]
		user, _ := s.DB.GetUser(context.Background(), username)

		if user.Name != "" {
			return fmt.Errorf("the username %v already exists", username)
		}
		id := uuid.New()
		created_at := time.Now()
		updated_at := time.Now()
		query_details_to_register := database.CreateUserParams{
			ID:        id,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			Name:      username,
		}
		newUser, err := s.DB.CreateUser(context.Background(), query_details_to_register)
		if err != nil {
			return err
		} else {
			err := s.Config.SetUser(newUser.Name)
			if err != nil {
				return err
			} else {
				fmt.Println("user has been created")
			}

		}
	}

	return nil

}
func handlerReset(s *state.State, cmd Command) error {

	err := s.DB.DelUsers(context.Background())
	if err != nil {
		return err
	}

	return nil

}
func handlerUsers(s *state.State, cmd Command) error {

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	} else {
		const appendCurrent = "(current)"
		for user := range users {
			if s.Config.CurrentUserName == users[user].Name {
				fmt.Printf("%v %v\n", users[user].Name, appendCurrent)
			} else {
				fmt.Printf("%v \n", users[user].Name)
			}

		}

	}

	return nil

}
func handlerRSSFeed(s *state.State, cmd Command) error {
	const feedURL = "https://www.wagslane.dev/index.xml"
	rssResponse,err:=rss.FetchFeed(context.Background(),feedURL)
	if err != nil {
		return err
	}
	fmt.Println(rssResponse)

	return nil

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
	commands.register("agg",handlerRSSFeed)
	return commands
}
