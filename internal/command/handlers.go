package command

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Dhairya3124/blog-aggregator/internal/database"
	"github.com/Dhairya3124/blog-aggregator/internal/rss"
	"github.com/Dhairya3124/blog-aggregator/internal/state"
	"github.com/google/uuid"
)

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
		return fmt.Errorf("register expects a single argument")
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

func handlerAggregateRSSFeed(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0{
		return fmt.Errorf("not enough arguments provided")

	}
	rss.ScrapeFeeds(context.Background(),s.DB,cmd.Args[0])

	return nil

}
func handlerRSSFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("addfeed expects a single argument which is url")
	} else if len(cmd.Args) == 1 {
		return fmt.Errorf("it expects two arguments")
	} else {
		name := cmd.Args[0]
		url := cmd.Args[1]

		id := uuid.New()
		created_at := time.Now()
		updated_at := time.Now()
		query_for_creating_feed := database.CreateFeedParams{
			ID:        id,
			Name:      name,
			Url:       url,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			UserID:    user.ID,
		}

		newFeed, err := s.DB.CreateFeed(context.Background(), query_for_creating_feed)
		if err != nil {
			return err
		}
		fmt.Println(newFeed)
		query_for_creating_follow := database.CreateFeedFollowParams{
			ID:        id,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			FeedID:    newFeed.ID,
			UserID:    user.ID,
		}
		_, err = s.DB.CreateFeedFollow(context.Background(), query_for_creating_follow)
		if err != nil {
			return err
		}
	}
	return nil
}
func handlerShowRSSFeed(s *state.State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(feeds)
	return nil
}
func handlerFollowRSSFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("follow expects a url")
	}
	feedURL := cmd.Args[0]
	feed, err := s.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	} else {
		id := uuid.New()
		created_at := time.Now()
		updated_at := time.Now()
		userId := user.ID
		feedId := feed.ID
		query_for_creating_follow := database.CreateFeedFollowParams{
			ID:        id,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			FeedID:    feedId,
			UserID:    userId,
		}
		follow, err := s.DB.CreateFeedFollow(context.Background(), query_for_creating_follow)
		if err != nil {
			return err
		}
		fmt.Println(follow)
	}
	return nil
}
func handlerFollowingRSSFeed(s *state.State, cmd Command, user database.User) error {
	userId := user.ID
	follows, err := s.DB.GetFollowsForUser(context.Background(), userId)
	if err != nil {
		return err
	} else {
		fmt.Println(follows)
	}
	return nil
}
func handlerUnfollowRSSFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("unfollow expects a url")
	}else{
		feedURL:=cmd.Args[0]
		query_for_creating_unfolllow:=database.DeleteFeedFollowParams{
			UserID: user.ID,
			Url: feedURL,
		}
		err:=s.DB.DeleteFeedFollow(context.Background(),query_for_creating_unfolllow)
		if err != nil {
			return err
		}
	}

return nil
}
func handlerBrowsePosts(s *state.State, cmd Command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) != 0{
		limitArg,_ := strconv.Atoi(cmd.Args[0])
		limit = int32(limitArg)
	}
	posts,err:=s.DB.GetPostsForUser(context.Background(),database.GetPostsForUserParams{UserID: user.ID,Limit: limit,})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println("Title:", post.Title, "URL:", post.Url)
	}

return nil
}
