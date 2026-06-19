package main

import (
	"context"
	"fmt"
	"github.com/Edudlufetips1/Gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
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

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed> <user>\n")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	err = s.db.UnfollowFeedForUser(context.Background(), database.UnfollowFeedForUserParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed %s successfully unfollowed", feed.Url)
	return nil
}
