package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	PrintableTweet() string
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int

	SetUser(string)
	SetText(string)
	SetDate(*time.Time)
	SetId(int)
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

type ImageTweet struct {
	TextTweet
	Url string
}

type QuotedTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func (textTweet *TextTweet) GetDate() *time.Time {
	return textTweet.Date
}

func (textTweet *TextTweet) GetUser() string {
	return textTweet.User
}

func (textTweet *TextTweet) GetId() int {
	return textTweet.Id
}

func (textTweet *TextTweet) GetText() string {
	return textTweet.Text
}

func (imageTweet *ImageTweet) GetUrl() string {
	return imageTweet.Url
}

func (textTweet *TextTweet) SetUser(user string) {
	textTweet.User = user
}

func (textTweet *TextTweet) SetText(text string) {
	textTweet.Text = text
}

func (textTweet *TextTweet) SetDate(date *time.Time) {
	textTweet.Date = date
}

func (textTweet *TextTweet) SetId(id int) {
	textTweet.Id = id
}

func NewTweet(user, text string) *TextTweet {
	return &TextTweet{User: user, Text: text}
}

func (textTweet *TextTweet) PrintableTweet() string {
	return "@" + textTweet.User + ": " + textTweet.Text
}

func (textTweet *TextTweet) String() string {
	return textTweet.PrintableTweet()
}

func NewImageTweet(user, text, url string) *ImageTweet {
	return &ImageTweet{TextTweet: TextTweet{User: user, Text: text}, Url: url}
}

func NewQuotedTweet(user, text string, quotedTweet Tweet) *QuotedTweet {
	return &QuotedTweet{TextTweet: TextTweet{User: user, Text: text},
		QuotedTweet: quotedTweet}
}

//func (quotedTweet *QuotedTweet) GetQuotedTweet()

func (imageTweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s\n%s", imageTweet.GetUser(), imageTweet.GetText(), imageTweet.GetUrl())
}

func (quotedTweet *QuotedTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%s\"", quotedTweet.GetUser(), quotedTweet.GetText(), quotedTweet.QuotedTweet.PrintableTweet())
}
