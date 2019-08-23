package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.TextTweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}

	//validation con Testify
	assert.Error(t, err)

	assert.Equal(t, "user is required", err.Error())
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	var user = "pepito"
	var text string
	tweet = domain.NewTweet(user, text)

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//validation con Testify
	assert.Error(t, err)

	assert.Equal(t, "text is required", err.Error())
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	user := "pepito"
	var text = "texto demasiado largo......................................................................................................................................................................................................................."
	tweet = domain.NewTweet(user, text)

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//validation con Testify
	assert.Error(t, err)

	assert.Equal(t, "text is too long", err.Error())
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	//Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet *domain.TextTweet

	//Fill the tweets with data
	tweet = domain.NewTweet("usuario uno", "texto uno")
	secondTweet = domain.NewTweet("usuario dos", "texto dos")

	//Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	//Validation
	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, 0, "usuario uno", "texto uno")
	isValidTweet(t, secondPublishedTweet, 1, "usuario dos", "texto dos")

}

//isValidTweet validates tweet
func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) {
	assert.Equal(t, user, tweet.GetUser())
	assert.Equal(t, text, tweet.GetText())
	assert.Equal(t, id, tweet.GetId())
}

func TestCanRetrieveTweetById(t *testing.T) {
	//initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet
	var id int
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	id, _ = tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	// Operation
	count := tweetManager.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	assert.Len(t, tweets, 2)

}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)
	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	tweetWriter.Write(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("...")
	}

	if savedTweet.GetId() != id {
		t.Errorf("...")
	}

}
