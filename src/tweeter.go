package main

import (
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/rest"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

func main() {

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)
	// Operation
	tweetManager.PublishTweet(tweet)

	ginServer := rest.NewGinServer(tweetManager)
	ginServer.StartServer()
	/*
		shell := ishell.New()
		shell.SetPrompt("Tweeter >> ")
		shell.Print("Type 'help' to know commands\n")

		shell.AddCmd(&ishell.Cmd{
			Name: "publishTweet",
			Help: "Publishes a tweet",
			Func: func(c *ishell.Context) {

				defer c.ShowPrompt(true)

				c.Print("Write your user: ")

				user := c.ReadLine()

				c.Print("Write your tweet: ")

				text := c.ReadLine()

				tweet := domain.NewTweet(user, text)

				service.PublishTweet(tweet)

				c.Print("Tweet sent\n")

				return
			},
		})

		shell.AddCmd(&ishell.Cmd{
			Name: "showTweet",
			Help: "Shows a tweet",
			Func: func(c *ishell.Context) {

				defer c.ShowPrompt(true)

				c.Print("Write twit ID: ")

				id, _ := strconv.Atoi(c.ReadLine())

				tweet := service.GetTweetById(id)

				c.Println(tweet)

				return
			},
		})

		shell.Run()
	*/
}
