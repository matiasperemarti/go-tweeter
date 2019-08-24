package service

import (
	"os"

	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
)

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWriter struct {
	tweet domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

const tweetFilePath = "tweets.txt"

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		tweetFilePath,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
	return &FileTweetWriter{file: file}
}

func (memoryTweetWriter *MemoryTweetWriter) Write(tweetToWrite domain.Tweet) {
	memoryTweetWriter.tweet = tweetToWrite
}

func (memoryTweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return memoryTweetWriter.tweet
}

func (ftw *FileTweetWriter) Write(tweet domain.Tweet) {
	go func() {
		ftw.file.WriteString(tweet.PrintableTweet() + "\n")
	}()
}
