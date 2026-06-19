package main

import (
	"context"
	"fmt"
	"time"
	"github.com/Edudlufetips1/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %s\n", feedFollow.FeedName)
	fmt.Printf("Current user: %s\n", feedFollow.UserName)
	return nil
}

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("'Following' command takes no arguments")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, ff := range feedFollows {
		fmt.Printf("Currently followed feeds by %s: %s\n", user.Name, ff.FeedName)
	}
	return nil
}