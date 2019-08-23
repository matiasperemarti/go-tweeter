package service

import (
	"fmt"
	"time"

	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
)

type TweetManager struct {
	tweets       []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
	tweetWriter  TweetWriter
}

//InitializeService publishes tweet
func NewTweetManager(tweetWriterIn TweetWriter) *TweetManager {
	tweetManager := &TweetManager{
		tweets:       make([]domain.Tweet, 0),
		tweetsByUser: make(map[string][]domain.Tweet),
		tweetWriter:  tweetWriterIn,
	}
	return tweetManager
}

//func InitializeService() {
//	tweets = make([]domain.Tweet, 0)
//	tweetsByUser = make(map[string][]domain.Tweet)
//}

// GetTweets gets tweets
func (tweetManager *TweetManager) GetTweets() []domain.Tweet {
	return tweetManager.tweets
}

//GetTweetByYd gets tweet by id
func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet {
	return tweetManager.tweets[id]
}

//PublishTweet publishes tweet
func (tweetManager *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (id int, err error) {
	err = validateTweetBeforePublishing(tweetToPublish)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	tweetToPublish.SetDate(&now)

	tweetToPublish.SetId(len(tweetManager.tweets))

	tweetManager.tweets = append(tweetManager.tweets, tweetToPublish)

	if _, exists := tweetManager.tweetsByUser[tweetToPublish.GetUser()]; !exists {
		tweetManager.tweetsByUser[tweetToPublish.GetUser()] = make([]domain.Tweet, 0)
	}

	tweetManager.tweetsByUser[tweetToPublish.GetUser()] = append(tweetManager.tweetsByUser[tweetToPublish.GetUser()], tweetToPublish)

	return tweetToPublish.GetId(), nil
}

//isValidTweet validates tweet
func validateTweetBeforePublishing(tweet domain.Tweet) error {
	if tweet.GetUser() == "" {
		return fmt.Errorf("user is required")
	}
	if tweet.GetText() == "" {
		return fmt.Errorf("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return fmt.Errorf("text is too long")
	}
	return nil
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	var result int

	for _, tweet := range tweetManager.tweets {
		if tweet.GetUser() == user {
			result++
		}
	}
	return result
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tweetManager.tweetsByUser[user]
}
