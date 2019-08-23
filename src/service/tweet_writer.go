package service

import "gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWriter struct {
	tweet domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

func (memoryTweetWriter *MemoryTweetWriter) Write(tweetToWrite domain.Tweet) {
	memoryTweetWriter.tweet = tweetToWrite
}

func (memoryTweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return memoryTweetWriter.tweet
}
