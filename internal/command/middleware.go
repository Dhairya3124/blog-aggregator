package command

import (
	"context"
	"fmt"

	"github.com/Dhairya3124/blog-aggregator/internal/database"
	"github.com/Dhairya3124/blog-aggregator/internal/state"
)


func middlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func( *state.State,  Command) error{
	return func(s *state.State, cmd Command) error {

		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not found due to the following error: %v", err)
		}

		return handler(s, cmd, user)
	}

}